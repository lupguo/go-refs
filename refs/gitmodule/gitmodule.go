// Package gitmodule 相关代码
// 包含 gitmodule 解析.
package gitmodule

import (
"bufio"
"bytes"
"errors"
"fmt"
"regexp"
"strings"
)

// Submodule gitmodule 信息.
type Submodule struct {
	Name string `json:"name"`
	Path string `json:"path"`
	URL  string `json:"url"`
}

const (
	// .gitmodules 中 gitmodule 在主仓中的路径.
	pathKey = "path"
	// .gitmodules 中 gitmodule 仓库地址.
	urlKey = "url"
	// 匹配出 gitmodule name 正则所用的 pattern.
	nameRegPattern = `^\[gitmodule\s+"(.*)"\]`

	// 匹配 http 格式的 git 仓库地址
	gitHttpURLRegPattern = `^(http|https)://(.*?)/(.*)\.git`
	// 匹配 ssh 格式的 git 仓库地址
	gitSSHURLRegPattern = `^git@(.*):(.*)\.git`
)

// ParseSubmodules 解析 .gitmodules 文件
// 每一行读取，遇到 '[' 说明为新一个 gitmodule 的开始
// 此时校验并保存上一个已经解析完成的 gitmodule.
func ParseSubmodules(gitmodules []byte) ([]*Submodule, error) {
	cfg := bufio.NewScanner(bytes.NewReader(gitmodules))
	var modules []*Submodule
	var submodule *Submodule
	var err error
	// key:value --> sshURL:originURL
	existRepos := make(map[string]string)
	for cfg.Scan() {
		line, isValid := checkAndTrimLine(cfg.Text())
		if !isValid {
			continue
		}
		if strings.HasPrefix(line, "[") {
			err := checkPreviousModuleFilled(submodule)
			if err != nil {
				return nil, err
			}
			name, err := parseModuleNameFromLine(line)
			if err != nil {
				return nil, err
			}
			submodule = &Submodule{Name: name}
			modules = append(modules, submodule)
			continue
		}
		existRepos, err = parseValue(line, existRepos, submodule)
		if err != nil {
			return nil, err
		}
	}
	return modules, nil
}

// ConvertURLToSSH 将 git http url 转换成 ssh url
// 如果输入是 ssh url 的话直接返回.
func ConvertURLToSSH(url string) (string, error) {
	httpPattern := regexp.MustCompile(gitHttpURLRegPattern)
	sshPattern := regexp.MustCompile(gitSSHURLRegPattern)
	if httpPattern.MatchString(url) {
		result := httpPattern.FindAllStringSubmatch(url, -1)
		// 由于条件中已经进行了 match 判定，所以这里可以认为肯定有值，不会产生数组越界
		url = fmt.Sprintf("git@%s:%s.git", result[0][2], result[0][3])
		return url, nil
	}
	if sshPattern.MatchString(url) {
		return url, nil
	}
	return "", errors.New(fmt.Sprintf("illegal git url: %s", url))
}

func parseValue(line string, existRepos map[string]string, module *Submodule) (map[string]string, error) {
	pattern := regexp.MustCompile(`(.*)=(.*)`)
	result := pattern.FindAllStringSubmatch(line, -1)
	if len(result) == 0 {
		return existRepos, nil
	}
	if len(result[0]) < 3 {
		return existRepos, nil
	}
	key := strings.TrimSpace(result[0][1])
	value := strings.TrimSpace(result[0][2])
	var err error
	// url 需要进行是否重复的校验以及统一转换成ssh格式
	if key == urlKey {
		originURL := value
		value, err = ConvertURLToSSH(value)
		if err != nil {
			return nil, err
		}
		existRepos, err = isRepoURLDuplicated(existRepos, originURL, value)
		if err != nil {
			return nil, err
		}
	}
	setValue(key, value, module)
	return existRepos, nil
}

func setValue(key, value string, module *Submodule) {
	if key == pathKey {
		module.Path = value
	}
	if key == urlKey {
		module.URL = value
	}
}

// checkAndTrimLine 校验并格式化当前行，不能为注释或者空行.
func checkAndTrimLine(line string) (string, bool) {

	if strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";") {
		return line, false
	}
	if index := strings.Index(line, "#"); index > 0 {
		line = line[:index]
	}
	if index := strings.Index(line, ";"); index > 0 {
		line = line[:index]
	}
	line = strings.TrimSpace(line)
	if len(line) == 0 {
		return line, false
	}
	return line, true
}

func checkPreviousModuleFilled(previousModule *Submodule) error {
	if previousModule == nil {
		return nil
	}
	if previousModule.Name == "" {
		return errors.New("exist empty gitmodule name")
	}
	if previousModule.Path == "" {
		return errors.New(fmt.Sprintf("gitmodule %s path empty", previousModule.Name))
	}
	if previousModule.URL == "" {
		return errors.New(fmt.Sprintf("gitmodule %s url empty", previousModule.Name))
	}
	return nil
}

func parseModuleNameFromLine(sectionName string) (string, error) {
	pattern := regexp.MustCompile(nameRegPattern)
	result := pattern.FindAllStringSubmatch(sectionName, -1)
	if len(result) == 0 || len(result[0]) < 2 {
		return "", fmt.Errorf("failed to match gitmodule name")
	}
	return result[0][1], nil
}

func isRepoURLDuplicated(existRepos map[string]string, originURL, sshURL string) (map[string]string, error) {
	existURL, exist := existRepos[sshURL]
	if exist {
		return nil, errors.New(fmt.Sprintf("git url %s and %s duplicated", existURL, originURL))
	}
	existRepos[sshURL] = originURL
	return existRepos, nil
}
