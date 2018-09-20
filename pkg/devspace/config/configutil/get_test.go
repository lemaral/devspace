package configutil

import (
	"testing"

	"github.com/covexo/devspace/pkg/devspace/config/v1"
	"github.com/juju/errors"
	"github.com/stretchr/testify/assert"
)

func TestGetConfigGlobalOnly(t *testing.T) {

	var err error
	config, err = makeGlobalConfig()
	if err != nil {
		t.Error(errors.Trace(err))
		return
	}

	defer func() {
		config = makeConfig()
		configRaw = makeConfig()
		overwriteConfig = makeConfig()
		overwriteConfigRaw = makeConfig()
	}()

	SaveConfig()

	configFromGet := GetConfig(true)

	assert.Equal(t, *config.Version, *configFromGet.Version)
	assert.Equal(t, *config.Cluster.UseKubeConfig, *configFromGet.Cluster.UseKubeConfig)
	assert.Equal(t, (*(*config.DevSpace.PortForwarding)[0].LabelSelector)["globalConfigMapKey"], (*(*configFromGet.DevSpace.PortForwarding)[0].LabelSelector)["globalConfigMapKey"])
	assert.Equal(t, (*(*config.DevSpace.PortForwarding)[0].PortMappings)[0].LocalPort, (*(*configFromGet.DevSpace.PortForwarding)[0].PortMappings)[0].LocalPort)
	assert.Equal(t, (*(*config.DevSpace.PortForwarding)[0].PortMappings)[0].RemotePort, (*(*configFromGet.DevSpace.PortForwarding)[0].PortMappings)[0].RemotePort)
	assert.Equal(t, *(*config.DevSpace.Sync)[0].ExcludePaths, *(*configFromGet.DevSpace.Sync)[0].ExcludePaths)
	assert.Equal(t, (*config.DevSpace.Release.Values)["key"], (*configFromGet.DevSpace.Release.Values)["key"])
	assert.Equal(t, (*config.DevSpace.Release.Values)[0], (*configFromGet.DevSpace.Release.Values)[0])

}

func TestGetConfigOverwriteOnly(t *testing.T) {

	var err error
	config, err = makeGlobalConfig()
	if err != nil {
		t.Error(errors.Trace(err))
		return
	}
	overwriteConfig, err = makeOverwriteConfig()
	if err != nil {
		t.Error(errors.Trace(err))
		return
	}

	defer func() {
		config = makeConfig()
		configRaw = makeConfig()
		overwriteConfig = makeConfig()
		overwriteConfigRaw = makeConfig()
	}()

	SaveConfig()

	configFromGet := GetConfig(true)

	assert.Equal(t, *overwriteConfig.Version, *configFromGet.Version)
	assert.Equal(t, *overwriteConfig.Cluster.UseKubeConfig, *configFromGet.Cluster.UseKubeConfig)
	assert.Equal(t, (*(*overwriteConfig.DevSpace.PortForwarding)[0].LabelSelector)["globalConfigMapKey"], (*(*configFromGet.DevSpace.PortForwarding)[0].LabelSelector)["globalConfigMapKey"])
	assert.Equal(t, (*(*overwriteConfig.DevSpace.PortForwarding)[0].PortMappings)[0].LocalPort, (*(*configFromGet.DevSpace.PortForwarding)[0].PortMappings)[0].LocalPort)
	assert.Equal(t, (*(*overwriteConfig.DevSpace.PortForwarding)[0].PortMappings)[0].RemotePort, (*(*configFromGet.DevSpace.PortForwarding)[0].PortMappings)[0].RemotePort)
	assert.Equal(t, *(*overwriteConfig.DevSpace.Sync)[0].ExcludePaths, *(*configFromGet.DevSpace.Sync)[0].ExcludePaths)
	assert.Equal(t, (*overwriteConfig.DevSpace.Release.Values)["key"], (*configFromGet.DevSpace.Release.Values)["key"])
	//assert.Equal(t, (*overwriteConfig.DevSpace.Release.Values)[0], (*configFromGet.DevSpace.Release.Values)[0])

}

func makeGlobalConfig() (*v1.Config, error) {

	globalConfig := makeConfig()

	if globalConfig == nil {
		return nil, errors.New("globalConfig is nil")
	}

	globalConfig.Version = new(string)
	*globalConfig.Version = "globalConfigVersion"

	globalConfig.Cluster.UseKubeConfig = new(bool)
	*globalConfig.Cluster.UseKubeConfig = true

	testPortForwarding := &v1.PortForwardingConfig{}
	testPortForwarding.LabelSelector = new(map[string]*string)
	*testPortForwarding.LabelSelector = make(map[string]*string)
	(*testPortForwarding.LabelSelector)["globalConfigMapKey"] = new(string)
	*(*testPortForwarding.LabelSelector)["globalConfigMapKey"] = "globalConfigMapValue"

	testPortMapping := &v1.PortMapping{}
	testPortMapping.LocalPort = new(int)
	*testPortMapping.LocalPort = 1
	testPortMapping.RemotePort = new(int)
	*testPortMapping.RemotePort = 1
	testPortForwarding.PortMappings = &[]*v1.PortMapping{
		testPortMapping,
	}
	*globalConfig.DevSpace.PortForwarding = []*v1.PortForwardingConfig{
		testPortForwarding,
	}

	testSyncConfig := &v1.SyncConfig{}
	testSyncConfig.ExcludePaths = new([]string)
	*testSyncConfig.ExcludePaths = []string{
		"globalConfigExcludePath",
	}
	globalConfig.DevSpace.Sync = new([]*v1.SyncConfig)
	*globalConfig.DevSpace.Sync = []*v1.SyncConfig{
		testSyncConfig,
	}

	globalConfig.DevSpace.Release.Values = new(map[interface{}]interface{})
	*globalConfig.DevSpace.Release.Values = make(map[interface{}]interface{})
	(*globalConfig.DevSpace.Release.Values)["key"] = "globalConfigInterface"
	(*globalConfig.DevSpace.Release.Values)[0] = "globalConfigInterface"

	return globalConfig, nil

}

func makeOverwriteConfig() (*v1.Config, error) {

	localConfig := makeConfig()

	if localConfig == nil {
		return nil, errors.New("localConfig is nil")
	}

	localConfig.Version = new(string)
	*localConfig.Version = "localConfigVersion"

	localConfig.Cluster.UseKubeConfig = new(bool)
	*localConfig.Cluster.UseKubeConfig = false

	testPortForwarding := &v1.PortForwardingConfig{}
	testPortForwarding.LabelSelector = new(map[string]*string)
	*testPortForwarding.LabelSelector = make(map[string]*string)
	(*testPortForwarding.LabelSelector)["localConfigMapKey"] = new(string)
	*(*testPortForwarding.LabelSelector)["localConfigMapKey"] = "localConfigMapValue"

	testPortMapping := &v1.PortMapping{}
	testPortMapping.LocalPort = new(int)
	*testPortMapping.LocalPort = 2
	testPortMapping.RemotePort = new(int)
	*testPortMapping.RemotePort = 2
	testPortForwarding.PortMappings = &[]*v1.PortMapping{
		testPortMapping,
	}
	*localConfig.DevSpace.PortForwarding = []*v1.PortForwardingConfig{
		testPortForwarding,
	}

	testSyncConfig := &v1.SyncConfig{}
	testSyncConfig.ExcludePaths = new([]string)
	*testSyncConfig.ExcludePaths = []string{
		"localConfigExcludePath",
	}
	localConfig.DevSpace.Sync = new([]*v1.SyncConfig)
	*localConfig.DevSpace.Sync = []*v1.SyncConfig{
		testSyncConfig,
	}

	localConfig.DevSpace.Release.Values = new(map[interface{}]interface{})
	*localConfig.DevSpace.Release.Values = make(map[interface{}]interface{})
	(*localConfig.DevSpace.Release.Values)["key"] = "localConfigInterface"
	(*localConfig.DevSpace.Release.Values)[0] = "localConfigInterface"

	return localConfig, nil

}
