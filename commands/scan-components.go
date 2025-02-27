package commands

import (
	"errors"
	"github.com/jfrog/jfrog-cli-core/plugins/components"
	"github.com/jfrog/jfrog-cli-plugin-template/scanUtils"
	"strconv"
	//"time"
)

func ScanComponents() components.Command {
	return components.Command{
		Name:        "scan-components",
		Description: "Scans a list of Packages/Components using Xray",
		Aliases:     []string{"sc"},
		Arguments:   getScanComponentsArguments(),
		Flags:       getScanComponentsFlags(),
		Action: func(c *components.Context) error {
			return scanComponents(c)
		},
	}
}

func getScanComponentsFlags() []components.Flag {
	return []components.Flag{
		components.StringFlag{
			Name: "v",
			Description: "\"high\" If you need only high vulnernability information " +
				"\"all\" for all the vulnerability information",
		},
		components.StringFlag{
			Name:        "l",
			Description: "To fetch all the license information ",
		},
	}
}

func getScanArguments() []components.Argument {
	return []components.Argument{
		{
			Name:        "Component Name",
			Description: "Name of the components(String of arrays) which Xray has to scan",
		},
	}
}

func getScanComponentsArguments() []components.Argument {
	return []components.Argument{
		{
			Name:        "Component Name",
			Description: "Name of the component which Xray has to scan",
		},
	}
}

func scanComponents(c *components.Context) error {
	if len(c.Arguments) == 0 {
		return errors.New("Wrong number of arguments. Expected: String array, " + "Received: " + strconv.Itoa(len(c.Arguments)))
	}
	compNames := c.Arguments
	var conf = new(scanUtils.ScanConfiguration)
	conf.VulnFlag = c.GetStringFlagValue("v")
	conf.LicenseFlag = c.GetStringFlagValue("l")

	rtDetails, err := scanUtils.GetRtDetails(c)
	if err != nil {
		return err
	}

	return scanUtils.ScanPackages(compNames, conf, rtDetails)
}
