package engine

import "testing"

func TestKeypressTask_GenerateYamlFile(t *testing.T) {
	t.Error(NewKeypressTaskDefault().GenerateYamlFile("./conf.yaml"))
}
