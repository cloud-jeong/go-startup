package main

import (
	"log"

	"k8s.io/helm/pkg/helm"
	"fmt"
	"k8s.io/helm/pkg/timeconv"
	"k8s.io/helm/pkg/proto/hapi/release"

	"encoding/json"
	"gopkg.in/yaml.v2"
	"strings"
)

const (
	tillerHost = "35.200.35.0:31904"
)

type listResult struct {
	Next     string
	Releases []listRelease
}

type listRelease struct {
	Name       string
	Revision   int32
	Updated    string
	Status     string
	Chart      string
	AppVersion string
	Namespace  string
}


func main() {
	helmClient := helm.NewClient(helm.Host(tillerHost))

	err := helmClient.PingTiller()
	if err != nil {
		log.Fatalf("tiller ping error: %v", err)
	}

	res, err := helmClient.GetReleasesListReleases()

	if err != nil {
		log.Fatalf("err: %v", err)
	}

	//for release := range res.GetReleases() {
	//	fmt.Printf("%v\n", release)
	//}

	//for _, r := range res.GetReleases() {
	//	name, version := r.GetName(), r.GetVersion()
	//
	//	fmt.Printf("%s, %d\n", name, version)
	//}

	result := getListResult(res.GetReleases(), res.Next)

	output, err := formatResult(l.output, l.short, result, l.colWidth)

	fmt.Fprintln(l.out, output)


	//fmt.Printf("%v\n", res.GetReleases())
}

func getListResult(rels []*release.Release, next string) listResult {
	listReleases := []listRelease{}
	for _, r := range rels {
		md := r.GetChart().GetMetadata()
		t := "-"
		if tspb := r.GetInfo().GetLastDeployed(); tspb != nil {
			t = timeconv.String(tspb)
		}

		lr := listRelease{
			Name:       r.GetName(),
			Revision:   r.GetVersion(),
			Updated:    t,
			Status:     r.GetInfo().GetStatus().GetCode().String(),
			Chart:      fmt.Sprintf("%s-%s", md.GetName(), md.GetVersion()),
			AppVersion: md.GetAppVersion(),
			Namespace:  r.GetNamespace(),
		}
		listReleases = append(listReleases, lr)
	}

	return listResult{
		Releases: listReleases,
		Next:     next,
	}
}

func formatResult(format string, short bool, result listResult, colWidth uint) (string, error) {
	var output string
	var err error

	var shortResult []string
	var finalResult interface{}
	if short {
		shortResult = shortenListResult(result)
		finalResult = shortResult
	} else {
		finalResult = result
	}

	switch format {
	case "":
		if short {
			output = formatTextShort(shortResult)
		} else {
			output = formatText(result, colWidth)
		}
	case "json":
		o, e := json.Marshal(finalResult)
		if e != nil {
			err = fmt.Errorf("Failed to Marshal JSON output: %s", e)
		} else {
			output = string(o)
		}
	case "yaml":
		o, e := yaml.Marshal(finalResult)
		if e != nil {
			err = fmt.Errorf("Failed to Marshal YAML output: %s", e)
		} else {
			output = string(o)
		}
	default:
		err = fmt.Errorf("Unknown output format \"%s\"", format)
	}
	return output, err
}

func formatText(result listResult, colWidth uint) string {
	nextOutput := ""
	if result.Next != "" {
		nextOutput = fmt.Sprintf("\tnext: %s\n", result.Next)
	}

	//table := uitable.New()
	//table.MaxColWidth = colWidth
	//table.AddRow("NAME", "REVISION", "UPDATED", "STATUS", "CHART", "APP VERSION", "NAMESPACE")
	//for _, lr := range result.Releases {
	//	table.AddRow(lr.Name, lr.Revision, lr.Updated, lr.Status, lr.Chart, lr.AppVersion, lr.Namespace)
	//}

	return fmt.Sprintf("%s%s", nextOutput, "")
}

func formatTextShort(shortResult []string) string {
	return strings.Join(shortResult, "\n")
}

func shortenListResult(result listResult) []string {
	names := []string{}
	for _, r := range result.Releases {
		names = append(names, r.Name)
	}

	return names
}



