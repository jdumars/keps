package sigs

const ListTemplate = `
package sigs

// THIS IS AN AUTOGENERATED FILE, DO NOT EDIT BY HAND
// this package supports exported functions on SIGs and their associated
// subprojects and is generated by parsing SIG information at
//
// https://raw.githubusercontent.com/kubernetes/community/master/sigs.yaml
//
// this file SHOULD NOT export any methods as its intention is simply to
// provide data for eventually exported functions

// generatedSIGSet is the set of SIGs generated by KEP tooling helpers
// DO NOT EDIT BY HAND
var generatedSIGSet = map[SIG]bool{
	{{ with .SIGs }}
		{{ range . }}
			.Name: true,
		{{ end }}
	{{ end }}
}

// generatedSIGSubprojectMapping groups subprojects by their owning SIG
// DO NOT EDIT BY HAND
var generatedSIGSubprojectMapping = map[SIG]map[Subproject]bool{
	{{ with .SIGs }}
		{{ range . }}
			.Name: true,
		{{ end }}
	{{ end }}

}

var generatedSIGList = []SIG{
	{{ with .SIGs }}
		{{ range . }}
			.Name: map[string]bool {
				{{ with .Subprojects }}
					{{ range . }}
						.Name: true,
					{{ end }}
				{{ end }}
			},
		{{ end }}
	{{ end }}
}
`
