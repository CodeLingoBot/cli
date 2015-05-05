package domains

import (
	"github.com/Scalingo/cli/Godeps/_workspace/src/gopkg.in/errgo.v1"
	"github.com/Scalingo/cli/api"
	"github.com/Scalingo/cli/io"
)

func Remove(app string, domain string) error {
	d, err := findDomain(app, domain)
	if err != nil {
		return errgo.Mask(err)
	}

	err = api.DomainsRemove(app, d.ID)
	if err != nil {
		return errgo.Mask(err)
	}

	io.Status("The domain", d.Name, "has been deleted")
	return nil
}

func findDomain(app string, domain string) (api.Domain, error) {
	domains, err := api.DomainsList(app)
	if err != nil {
		return api.Domain{}, errgo.Mask(err)
	}

	for _, d := range domains {
		if d.Name == domain {
			return d, nil
		}
	}
	return api.Domain{}, errgo.New("There is no such domain, please ensure you've added it correctly.\nhttps://my.scalingo.com/apps/" + app + "/domains")
}
