module github.com/kevinrizza/offline-cataloger

go 1.12

require (
	github.com/ghodss/yaml v1.0.0
	github.com/go-openapi/analysis v0.19.0 // indirect
	github.com/go-openapi/errors v0.19.0 // indirect
	github.com/go-openapi/runtime v0.19.0
	github.com/go-openapi/strfmt v0.19.0
	github.com/go-openapi/swag v0.19.0 // indirect
	github.com/go-openapi/validate v0.19.0 // indirect
	github.com/golang/mock v1.2.1-0.20190329180013-73dc87cad333
	github.com/googleapis/gnostic v0.2.0 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/magiconair/properties v1.8.1 // indirect
	github.com/mailru/easyjson v0.0.0-20190403194419-1ea4449da983 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/operator-framework/go-appr v0.0.0-20180917210448-f2aef88446f2
	github.com/operator-framework/operator-lifecycle-manager v0.0.0-20190522155654-7b2b397ad0cc
	github.com/pelletier/go-toml v1.4.0 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cobra v0.0.4
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/viper v1.4.0
	github.com/stretchr/testify v1.3.0
	golang.org/x/crypto v0.0.0-20190530122614-20be4c3c3ed5 // indirect
	golang.org/x/oauth2 v0.0.0-20190523182746-aaccbc9213b0 // indirect
	golang.org/x/sys v0.0.0-20190602015325-4c4f7f33c9ed // indirect
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/appengine v1.6.0 // indirect
	google.golang.org/genproto v0.0.0-20181016170114-94acd270e44e // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	k8s.io/api v0.0.0-20190313235455-40a48860b5ab // indirect
	k8s.io/apiextensions-apiserver v0.0.0-20190315093550-53c4693659ed
	k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d
)

replace (
	// pin kube dependencies to release-1.12 branch
	github.com/evanphx/json-patch => github.com/evanphx/json-patch v0.0.0-20190203023257-5858425f7550
	k8s.io/api => k8s.io/api v0.0.0-20181128191700-6db15a15d2d3
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190221101132-cda7b6cfba78
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190221084156-01f179d85dbc
	k8s.io/apiserver => k8s.io/apiserver v0.0.0-20190402012035-5e1c1f41ee34
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190228133956-77e032213d34
	k8s.io/code-generator => k8s.io/code-generator v0.0.0-20181128191024-b1289fc74931
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.0.0-20190221095344-e77f03c95d65
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20180711000925-0cf8f7e6ed1d
)
