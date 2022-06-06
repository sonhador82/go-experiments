package selectel

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"net"
	"os"
	//"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/utils/openstack/clientconfig"
)

type Config struct {
	Username string
	Password string
	ProjectDomainName string
	ProjectName string
}


type selectel struct {
	pClient  *gophercloud.ProviderClient
}

type ServerPair struct {
	Name string
	IP net.IP
}


func NewSelectel(config *Config) *selectel  {
	opts := &clientconfig.ClientOpts{
		AuthInfo: &clientconfig.AuthInfo{
			AuthURL:     "https://api.selvpc.ru/identity/v3",
			Username:    config.Username,
			Password:    config.Password,
			ProjectDomainName: config.ProjectDomainName,
			ProjectName: config.ProjectName,
		},
	}
	pClient, err := clientconfig.AuthenticatedClient(opts)
	if err != nil {
		panic(err)
	}
	s := new(selectel)
	s.pClient = pClient
	return s
}

func (s *selectel) GetServerList(netBlock net.IPNet) []ServerPair {
	fmt.Println("get server list")

	nova, err := openstack.NewComputeV2(s.pClient, gophercloud.EndpointOpts{Region: os.Getenv("OS_REGION_NAME")})
	if err != nil {
		panic(err)
	}

	allPages, err := servers.List(nova, servers.ListOpts{}).AllPages()
	if err != nil {
		panic(err)
	}

	allSrv, err := servers.ExtractServers(allPages)
	if err != nil {
		panic(err)
	}

	pairs := make([]ServerPair, 0, 10)
	// ссаные костыли, тут типа возвращаем пары ip/имя, если попадает в диапазон
	for _, srv := range allSrv {
		for _, networkData  := range srv.Addresses {
			netInterfaces := networkData.([]interface{})
			for _, netPortData := range netInterfaces {
				netPort := netPortData.(map[string]interface{})
				addr := net.ParseIP(netPort["addr"].(string))

				if netBlock.Contains(addr) {
					pairs = append(pairs, ServerPair{
						Name: srv.Name,
						IP: addr,
					})
				}
			}

		}
	}
	return pairs
}
