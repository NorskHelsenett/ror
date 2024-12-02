package viewsmodels

import (
	"strings"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
)

type PolicyreportGlobalQueryType string

const (
	PolicyreportGlobalQueryTypeUnknown PolicyreportGlobalQueryType = "Unknown"
	PolicyreportGlobalQueryTypeCluster PolicyreportGlobalQueryType = "Cluster"
	PolicyreportGlobalQueryTypePolicy  PolicyreportGlobalQueryType = "Policy"
)

type PolicyreportGlobal struct {
	Cluster string `json:"cluster"`
	Policy  string `json:"policy"`
	Fail    int    `json:"fail"`
	Pass    int    `json:"pass"`
}

type PolicyreportGlobalQuery struct {
	Type     PolicyreportGlobalQueryType `json:"type"`
	Internal bool                        `json:"internal"`
}

type PolicyreportView struct {
	Clusterid  string                      `json:"clusterid" validate:""`
	Namespaces []PolicyreportViewNamespace `json:"namespaces" validate:""`
	PolicyreportSummary
}

type PolicyreportViewNamespace struct {
	Name     string                     `json:"name" validate:""`
	Policies []PolicyreportViewPolicies `json:"policies" validate:""`
	PolicyreportSummary
}
type PolicyreportViewPolicies struct {
	Name    string                    `json:"name" validate:""`
	Reports []PolicyreportViewReports `json:"reports" validate:""`
	PolicyreportSummary
}
type PolicyreportViewReports struct {
	Uid        string `json:"uid" validate:""`
	Name       string `json:"name" validate:""`
	ApiVersion string `json:"apiversion" validate:""`
	Kind       string `json:"kind" validate:""`
	Result     string `json:"result" validate:""`
	Category   string `json:"category" validate:""`
	Message    string `json:"message" validate:""`
}

type PolicyreportSummary struct {
	Fail  int `json:"failed" validate:""`
	Pass  int `json:"passed" validate:""`
	Error int `json:"error" validate:""`
	Warn  int `json:"warning" validate:""`
	Skip  int `json:"skipped" validate:""`
	Total int `json:"total" validate:""`
}

func (pr *PolicyreportView) UpdateSummary() {
	for i, namespace := range pr.Namespaces {
		for _, policy := range namespace.Policies {
			namespace.Pass = namespace.Pass + policy.Pass
			namespace.Fail = namespace.Fail + policy.Fail
			namespace.Error = namespace.Error + policy.Error
			namespace.Warn = namespace.Warn + policy.Warn
			namespace.Skip = namespace.Skip + policy.Skip

		}
		pr.Namespaces[i].Pass = namespace.Pass
		pr.Namespaces[i].Fail = namespace.Fail
		pr.Namespaces[i].Error = namespace.Error
		pr.Namespaces[i].Warn = namespace.Warn
		pr.Namespaces[i].Skip = namespace.Skip
		pr.Namespaces[i].Total = namespace.Pass + namespace.Fail + namespace.Error + namespace.Warn + namespace.Skip

		pr.Pass = pr.Pass + namespace.Pass
		pr.Fail = pr.Fail + namespace.Fail
		pr.Error = pr.Error + namespace.Error
		pr.Warn = pr.Warn + namespace.Warn
		pr.Skip = pr.Skip + namespace.Skip
		pr.Total = pr.Pass + pr.Fail + pr.Error + pr.Warn + pr.Skip
	}
}

func (pr *PolicyreportView) ImportData(reports apiresourcecontracts.ResourceListPolicyreports) error {
	pr.Clusterid = reports.Owner.Subject
	for _, policyreport := range reports.Policyreports {
		if !strings.HasPrefix(policyreport.Metadata.Name, "polr") {
			namespace := pr.FindOrCreatePolicyreportNamespace(policyreport.Metadata.Namespace)
			_ = namespace.AddPolicyreport(policyreport)
		}
	}
	pr.UpdateSummary()
	return nil
}

func (pr *PolicyreportView) FindOrCreatePolicyreportNamespace(namespaceToReturn string) *PolicyreportViewNamespace {
	namespace := pr.ReturnPolicyreportNamespace(namespaceToReturn)
	if namespace == nil {
		newNamespace := PolicyreportViewNamespace{
			Name: namespaceToReturn,
		}
		pr.Namespaces = append(pr.Namespaces, newNamespace)
		namespace = pr.ReturnPolicyreportNamespace(namespaceToReturn)
	}

	return namespace
}

func (pr PolicyreportView) ReturnPolicyreportNamespace(namespaceToReturn string) *PolicyreportViewNamespace {
	idx := pr.ReturnPolicyreportNamespaceIdx(namespaceToReturn)
	if idx != -1 {
		return &pr.Namespaces[idx]
	}
	return nil
}

func (pr PolicyreportView) ReturnPolicyreportNamespaceIdx(namespaceToReturn string) int {
	if len(pr.Namespaces) > 0 {
		for idx, namespace := range pr.Namespaces {
			if namespace.Name == namespaceToReturn {
				return idx
			}
		}
	}
	return -1
}

func (prn *PolicyreportViewNamespace) AddPolicyreport(reports apiresourcecontracts.ResourcePolicyReport) error {
	for _, policyreport := range reports.Results {
		policies := prn.FindOrCreatePolicyreportReport(policyreport.Policy, reports.Summary)
		_ = policies.AddPolicyreportPolicy(policyreport)
	}
	return nil
}

func (prn *PolicyreportViewNamespace) FindOrCreatePolicyreportReport(policyToReturn string, summary apiresourcecontracts.ResourcePolicyReportSummary) *PolicyreportViewPolicies {
	namespace := prn.ReturnPolicyreportPolicy(policyToReturn)
	if namespace == nil {
		newPolicy := PolicyreportViewPolicies{
			Name: policyToReturn,
			PolicyreportSummary: PolicyreportSummary{
				Pass:  summary.Pass,
				Fail:  summary.Fail,
				Error: summary.Error,
				Warn:  summary.Warn,
				Skip:  summary.Skip,
				Total: summary.Pass + summary.Fail + summary.Error + summary.Warn + summary.Skip,
			},
		}
		prn.Policies = append(prn.Policies, newPolicy)
		namespace = prn.ReturnPolicyreportPolicy(policyToReturn)
	}

	return namespace
}

func (prn PolicyreportViewNamespace) ReturnPolicyreportPolicy(policyToReturn string) *PolicyreportViewPolicies {
	idx := prn.ReturnPolicyreportPolicyIdx(policyToReturn)
	if idx != -1 {
		return &prn.Policies[idx]
	}
	return nil
}

func (prn PolicyreportViewNamespace) ReturnPolicyreportPolicyIdx(policyToReturn string) int {
	if len(prn.Policies) > 0 {
		for idx, report := range prn.Policies {
			if report.Name == policyToReturn {
				return idx
			}
		}
	}
	return -1
}

func (prp *PolicyreportViewPolicies) AddPolicyreportPolicy(result apiresourcecontracts.ResourcePolicyReportResults) error {

	for _, resource := range result.Resources {
		policy := PolicyreportViewReports{
			Uid:        resource.Uid,
			Name:       resource.Name,
			Kind:       resource.Kind,
			ApiVersion: resource.ApiVersion,
			Result:     result.Result,
			Category:   result.Category,
			Message:    result.Message,
		}
		prp.Reports = append(prp.Reports, policy)
	}

	return nil
}
