module github.com/Tencent/bk-bcs/bcs-scenarios/bcs-gitops-cli

go 1.22.7

replace (
	github.com/Tencent/bk-bcs/bcs-scenarios/bcs-gitops-manager => ../bcs-gitops-manager
	github.com/Tencent/bk-bcs/bcs-scenarios/bcs-gitops-workflow => ../bcs-gitops-workflow
	github.com/Tencent/bk-bcs/bcs-scenarios/bcs-terraform-controller => ../bcs-terraform-controller
	github.com/argoproj/gitops-engine => github.com/argoproj/gitops-engine v0.7.1-0.20240714153147-adb68bcaab73
	github.com/chai2010/gettext-go => github.com/chai2010/gettext-go v1.0.3

	// https://github.com/golang/go/issues/33546#issuecomment-519656923
	github.com/go-check/check => github.com/go-check/check v0.0.0-20180628173108-788fd7840127

	github.com/grpc-ecosystem/grpc-gateway => github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/micro/go-micro => go-micro.dev/v4/api v1.18.0
	go.opentelemetry.io/otel => go.opentelemetry.io/otel v1.33.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace => go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.33.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc => go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.33.0
	go.opentelemetry.io/otel/metric => go.opentelemetry.io/otel/metric v1.33.0
	go.opentelemetry.io/otel/sdk => go.opentelemetry.io/otel/sdk v1.33.0
	go.opentelemetry.io/otel/trace => go.opentelemetry.io/otel/trace v1.33.0
	google.golang.org/grpc => google.golang.org/grpc v1.66.2

	// Avoid CVE-2022-3064
	gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.4.0

	// Avoid CVE-2022-28948
	gopkg.in/yaml.v3 => gopkg.in/yaml.v3 v3.0.1

	k8s.io/api => k8s.io/api v0.29.6
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.29.6
	k8s.io/apimachinery => k8s.io/apimachinery v0.29.6
	k8s.io/apiserver => k8s.io/apiserver v0.29.6
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.29.6
	k8s.io/client-go => k8s.io/client-go v0.29.6
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.29.6
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.29.6
	k8s.io/code-generator => k8s.io/code-generator v0.29.6
	k8s.io/component-base => k8s.io/component-base v0.29.6
	k8s.io/component-helpers => k8s.io/component-helpers v0.29.6
	k8s.io/controller-manager => k8s.io/controller-manager v0.29.6
	k8s.io/cri-api => k8s.io/cri-api v0.29.6
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.29.6
	k8s.io/dynamic-resource-allocation => k8s.io/dynamic-resource-allocation v0.29.6
	k8s.io/endpointslice => k8s.io/endpointslice v0.29.6
	k8s.io/kms => k8s.io/kms v0.29.6
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.29.6
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.29.6

	// https://github.com/kubernetes/kubernetes/issues/79384#issuecomment-505627280
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20241212222426-2c72e554b1e7
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.29.6
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.29.6
	k8s.io/kubectl => k8s.io/kubectl v0.29.6
	k8s.io/kubelet => k8s.io/kubelet v0.29.6
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.29.6
	k8s.io/metrics => k8s.io/metrics v0.29.6
	k8s.io/mount-utils => k8s.io/mount-utils v0.29.6
	k8s.io/pod-security-admission => k8s.io/pod-security-admission v0.29.6
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.29.6
	k8s.io/sample-cli-plugin => k8s.io/sample-cli-plugin v0.29.6
	k8s.io/sample-controller => k8s.io/sample-controller v0.29.6
	sigs.k8s.io/kustomize/api => sigs.k8s.io/kustomize/api v0.17.2
	sigs.k8s.io/kustomize/kyaml => sigs.k8s.io/kustomize/kyaml v0.17.1
	sigs.k8s.io/structured-merge-diff/v4 => sigs.k8s.io/structured-merge-diff/v4 v4.4.1

	github.com/go-telegram-bot-api/telegram-bot-api/v5 => github.com/OvyFlash/telegram-bot-api/v5 v5.0.0-20240108230938-63e5c59035bf
)

require (
	github.com/Tencent/bk-bcs/bcs-common v0.0.0-20240115064206-f3812beb640b
	github.com/Tencent/bk-bcs/bcs-scenarios/bcs-gitops-manager v0.0.0-20240124084046-a0156afd4af9
	github.com/Tencent/bk-bcs/bcs-scenarios/bcs-gitops-vaultplugin-server v0.0.0-20240220073632-303c7a871f3b
	github.com/Tencent/bk-bcs/bcs-scenarios/bcs-gitops-workflow v0.0.0-00010101000000-000000000000
	github.com/Tencent/bk-bcs/bcs-scenarios/bcs-terraform-controller v0.0.0-20240507114009-045488ee0cd5
	github.com/argoproj/argo-cd/v2 v2.12.8
	github.com/fatih/color v1.16.0
	github.com/itchyny/json2yaml v0.1.4
	github.com/moul/http2curl v1.0.0
	github.com/olekukonko/tablewriter v0.0.5
	github.com/pkg/errors v0.9.1
	github.com/sergi/go-diff v1.3.2-0.20230802210424-5b0b94c5c0d3
	github.com/spf13/cobra v1.8.1
	golang.org/x/net v0.32.0
	gopkg.in/yaml.v3 v3.0.1
	k8s.io/apimachinery v0.31.4
	k8s.io/cli-runtime v0.31.0
	k8s.io/kubectl v0.31.2
)

require (
	cloud.google.com/go/compute/metadata v0.5.0 // indirect
	code.gitea.io/sdk/gitea v0.18.0 // indirect
	dario.cat/mergo v1.0.0 // indirect
	github.com/Azure/go-ansiterm v0.0.0-20230124172434-306776ec8161 // indirect
	github.com/Knetic/govaluate v3.0.1-0.20171022003610-9aa49832a739+incompatible // indirect
	github.com/MakeNowJust/heredoc v1.0.0 // indirect
	github.com/Masterminds/goutils v1.1.1 // indirect
	github.com/Masterminds/semver/v3 v3.2.1 // indirect
	github.com/Masterminds/sprig/v3 v3.2.3 // indirect
	github.com/Microsoft/go-winio v0.6.1 // indirect
	github.com/PagerDuty/go-pagerduty v1.7.0 // indirect
	github.com/ProtonMail/go-crypto v1.0.0 // indirect
	github.com/RocketChat/Rocket.Chat.Go.SDK v0.0.0-20210112200207-10ab4d695d60 // indirect
	github.com/Tencent/bk-bcs/bcs-common/pkg/audit v0.0.0-20231114032042-90745b766c4b // indirect
	github.com/Tencent/bk-bcs/bcs-common/pkg/auth v0.0.0-20231130061921-aba183933987 // indirect
	github.com/Tencent/bk-bcs/bcs-common/pkg/bcsapiv4 v0.0.0-20231114032042-90745b766c4b // indirect
	github.com/Tencent/bk-bcs/bcs-common/pkg/otel v0.0.0-20231124101643-0d25188774c2 // indirect
	github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubernetes/common v0.0.0-20220330120237-0bbed74dcf6d // indirect
	github.com/Tencent/bk-bcs/bcs-services/cluster-resources v0.0.0-20231114032042-90745b766c4b // indirect
	github.com/Tencent/bk-bcs/bcs-services/pkg v0.0.0-20231114032042-90745b766c4b // indirect
	github.com/Tencent/bk-bcs/bcs-services/pkg/bcs-auth-v4 v0.0.0-20240102072303-4173c9729642 // indirect
	github.com/TencentBlueKing/bk-audit-go-sdk v0.0.5 // indirect
	github.com/TencentBlueKing/gopkg v1.1.0 // indirect
	github.com/TencentBlueKing/iam-go-sdk v0.1.5 // indirect
	github.com/TomOnTime/utfutil v0.0.0-20180511104225-09c41003ee1d // indirect
	github.com/agext/levenshtein v1.2.1 // indirect
	github.com/alicebob/gopher-json v0.0.0-20230218143504-906a9b012302 // indirect
	github.com/alicebob/miniredis/v2 v2.30.4 // indirect
	github.com/antonmedv/expr v1.15.2 // indirect
	github.com/apparentlymart/go-textseg/v15 v15.0.0 // indirect
	github.com/argoproj/gitops-engine v0.7.3 // indirect
	github.com/argoproj/notifications-engine v0.4.1-0.20240606074338-0802cd427621 // indirect
	github.com/argoproj/pkg v0.13.7-0.20230626144333-d56162821bd1 // indirect
	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2 // indirect
	github.com/avast/retry-go v3.0.0+incompatible // indirect
	github.com/aws/aws-sdk-go v1.50.8 // indirect
	github.com/aws/aws-sdk-go-v2 v1.24.1 // indirect
	github.com/aws/aws-sdk-go-v2/config v1.25.12 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.16.16 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.14.11 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.2.10 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.5.10 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.7.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.10.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.10.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/sqs v1.29.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.18.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.21.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.26.7 // indirect
	github.com/aws/smithy-go v1.19.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bitly/go-simplejson v0.5.0 // indirect
	github.com/blang/semver/v4 v4.0.0 // indirect
	github.com/bmatcuk/doublestar/v4 v4.6.0 // indirect
	github.com/bombsimon/logrusr/v2 v2.0.1 // indirect
	github.com/bradleyfalzon/ghinstallation/v2 v2.6.0 // indirect
	github.com/casbin/casbin/v2 v2.77.2 // indirect
	github.com/cenkalti/backoff/v3 v3.2.2 // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/chai2010/gettext-go v1.0.2 // indirect
	github.com/chainguard-dev/git-urls v1.0.2 // indirect
	github.com/cloudflare/circl v1.3.7 // indirect
	github.com/coreos/go-oidc/v3 v3.6.0 // indirect
	github.com/coreos/go-semver v0.3.1 // indirect
	github.com/coreos/go-systemd/v22 v22.5.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.4 // indirect
	github.com/cyphar/filepath-securejoin v0.2.4 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/daviddengcn/go-colortext v1.0.0 // indirect
	github.com/davidmz/go-pageant v1.0.2 // indirect
	github.com/desertbit/timer v0.0.0-20180107155436-c41aec40b27f // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/distribution/reference v0.5.0 // indirect
	github.com/dlclark/regexp2 v1.11.2 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/emicklei/go-restful/v3 v3.11.0 // indirect
	github.com/emirpasic/gods v1.18.1 // indirect
	github.com/envoyproxy/protoc-gen-validate v1.1.0 // indirect
	github.com/evanphx/json-patch v5.9.0+incompatible // indirect
	github.com/evanphx/json-patch/v5 v5.8.0 // indirect
	github.com/exponent-io/jsonpath v0.0.0-20151013193312-d6023ce2651d // indirect
	github.com/fatih/camelcase v1.0.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/fvbommel/sortorder v1.1.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.2 // indirect
	github.com/gfleury/go-bitbucket-v1 v0.0.0-20220301131131-8e7ed04b843e // indirect
	github.com/go-acme/lego/v4 v4.4.0 // indirect
	github.com/go-errors/errors v1.4.2 // indirect
	github.com/go-fed/httpsig v1.1.0 // indirect
	github.com/go-git/gcfg v1.5.1-0.20230307220236-3a3c6141e376 // indirect
	github.com/go-git/go-billy/v5 v5.5.0 // indirect
	github.com/go-git/go-git/v5 v5.12.0 // indirect
	github.com/go-jose/go-jose/v3 v3.0.3 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-micro/plugins/v4/registry/etcd v1.1.0 // indirect
	github.com/go-openapi/analysis v0.21.4 // indirect
	github.com/go-openapi/errors v0.20.3 // indirect
	github.com/go-openapi/jsonpointer v0.21.0 // indirect
	github.com/go-openapi/jsonreference v0.20.2 // indirect
	github.com/go-openapi/loads v0.21.2 // indirect
	github.com/go-openapi/runtime v0.26.0 // indirect
	github.com/go-openapi/spec v0.20.8 // indirect
	github.com/go-openapi/strfmt v0.21.7 // indirect
	github.com/go-openapi/swag v0.23.0 // indirect
	github.com/go-openapi/validate v0.22.1 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.14.0 // indirect
	github.com/go-playground/webhooks/v6 v6.3.0 // indirect
	github.com/go-redis/cache/v9 v9.0.0 // indirect
	github.com/go-resty/resty/v2 v2.7.0 // indirect
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/go-telegram-bot-api/telegram-bot-api/v5 v5.5.1 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gobwas/httphead v0.1.0 // indirect
	github.com/gobwas/pool v0.2.1 // indirect
	github.com/gobwas/ws v1.2.1 // indirect
	github.com/gogits/go-gogs-client v0.0.0-20200905025246-8bb8a50cb355 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt/v4 v4.5.0 // indirect
	github.com/golang-migrate/migrate/v4 v4.16.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/btree v1.1.2 // indirect
	github.com/google/gnostic-models v0.6.9 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/go-github/v35 v35.3.0 // indirect
	github.com/google/go-github/v41 v41.0.0 // indirect
	github.com/google/go-github/v53 v53.2.0 // indirect
	github.com/google/go-jsonnet v0.20.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/s2a-go v0.1.7 // indirect
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.2 // indirect
	github.com/googleapis/gax-go/v2 v2.12.0 // indirect
	github.com/gorilla/handlers v1.5.1 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/gosimple/slug v1.13.1 // indirect
	github.com/gosimple/unidecode v1.0.1 // indirect
	github.com/gregdel/pushover v1.2.1 // indirect
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0 // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.24.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.7 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/go-secure-stdlib/parseutil v0.1.7 // indirect
	github.com/hashicorp/go-secure-stdlib/strutil v0.1.2 // indirect
	github.com/hashicorp/go-sockaddr v1.0.4 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/hashicorp/hcl v1.0.1-vault-5 // indirect
	github.com/hashicorp/hcl/v2 v2.19.1 // indirect
	github.com/hashicorp/terraform-exec v0.20.0 // indirect
	github.com/hashicorp/terraform-json v0.19.0 // indirect
	github.com/hashicorp/vault/api v1.10.0 // indirect
	github.com/huandu/xstrings v1.3.3 // indirect
	github.com/imdario/mergo v0.3.16 // indirect
	github.com/improbable-eng/grpc-web v0.15.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/itchyny/gojq v0.12.13 // indirect
	github.com/itchyny/timefmt-go v0.1.5 // indirect
	github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 // indirect
	github.com/jeremywohl/flatten v1.0.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/juju/ratelimit v1.0.1 // indirect
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/kevinburke/ssh_config v1.2.0 // indirect
	github.com/klauspost/compress v1.17.0 // indirect
	github.com/ktrysmt/go-bitbucket v0.9.67 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/liggitt/tabwriter v0.0.0-20181228230101-89fcab3d43de // indirect
	github.com/lithammer/dedent v1.1.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattbaird/jsonpatch v0.0.0-20200820163806-098863c1fc24 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/microsoft/azure-devops-go-api/azuredevops v1.0.0-b5 // indirect
	github.com/miekg/dns v1.1.50 // indirect
	github.com/minio/blake2b-simd v0.0.0-20160723061019-3f5f724cb5b1 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/mitchellh/hashstructure v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/moby/spdystream v0.4.0 // indirect
	github.com/moby/term v0.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/monochromegane/go-gitignore v0.0.0-20200626010858-205db1a8cc00 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/mxk/go-flowrate v0.0.0-20140419014527-cca7078d478f // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.1.0-rc4 // indirect
	github.com/openzipkin/zipkin-go v0.3.0 // indirect
	github.com/opsgenie/opsgenie-go-sdk-v2 v1.0.5 // indirect
	github.com/oxtoacart/bpool v0.0.0-20190530202638-03653db5a59c // indirect
	github.com/parnurzeal/gorequest v0.2.16 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/pjbgf/sha1cd v0.3.0 // indirect
	github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring v0.69.1 // indirect
	github.com/prometheus-operator/prometheus-operator/pkg/client v0.69.1 // indirect
	github.com/prometheus/client_golang v1.19.1 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.55.0 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	github.com/r3labs/diff v1.1.0 // indirect
	github.com/redis/go-redis/v9 v9.0.5 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/rs/cors v1.9.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/ryanuber/go-glob v1.0.0 // indirect
	github.com/samuel/go-zookeeper v0.0.0-20201211165307-7117e9ea2414 // indirect
	github.com/shopspring/decimal v1.2.0 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/skeema/knownhosts v1.2.2 // indirect
	github.com/skratchdot/open-golang v0.0.0-20160302144031-75fb7ed4208c // indirect
	github.com/slack-go/slack v0.12.2 // indirect
	github.com/soheilhy/cmux v0.1.5 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/cast v1.6.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/tidwall/gjson v1.14.4 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/ugorji/go/codec v1.2.11 // indirect
	github.com/urfave/cli/v2 v2.3.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	github.com/vmihailenco/go-tinylfu v0.2.2 // indirect
	github.com/vmihailenco/msgpack/v5 v5.3.5 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/xanzy/go-gitlab v0.91.1 // indirect
	github.com/xanzy/ssh-agent v0.3.3 // indirect
	github.com/xlab/treeprint v1.2.0 // indirect
	github.com/yuin/gopher-lua v1.1.0 // indirect
	github.com/zclconf/go-cty v1.14.1 // indirect
	go-micro.dev/v4 v4.10.2 // indirect
	go.etcd.io/etcd/api/v3 v3.5.14 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.14 // indirect
	go.etcd.io/etcd/client/v3 v3.5.14 // indirect
	go.mongodb.org/mongo-driver v1.12.1 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.53.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.53.0 // indirect
	go.opentelemetry.io/otel v1.33.0 // indirect
	go.opentelemetry.io/otel/exporters/jaeger v1.3.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.33.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.27.0 // indirect
	go.opentelemetry.io/otel/exporters/zipkin v1.3.0 // indirect
	go.opentelemetry.io/otel/metric v1.33.0 // indirect
	go.opentelemetry.io/otel/sdk v1.33.0 // indirect
	go.opentelemetry.io/otel/trace v1.33.0 // indirect
	go.opentelemetry.io/proto/otlp v1.4.0 // indirect
	go.starlark.net v0.0.0-20230525235612-a134d8f9ddca // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.26.0 // indirect
	golang.org/x/crypto v0.30.0 // indirect
	golang.org/x/exp v0.0.0-20231006140011-7918f672742d // indirect
	golang.org/x/mod v0.20.0 // indirect
	golang.org/x/oauth2 v0.24.0 // indirect
	golang.org/x/sync v0.10.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/term v0.27.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	golang.org/x/tools v0.24.0 // indirect
	gomodules.xyz/envconfig v1.3.1-0.20190308184047-426f31af0d45 // indirect
	gomodules.xyz/notify v0.1.1 // indirect
	google.golang.org/api v0.155.0 // indirect
	google.golang.org/genproto v0.0.0-20240227224415-6ceb2ff114de // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20241209162323-e6fa225c2576 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241209162323-e6fa225c2576 // indirect
	google.golang.org/grpc v1.68.1 // indirect
	google.golang.org/protobuf v1.35.2 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/evanphx/json-patch.v4 v4.12.0 // indirect
	gopkg.in/go-playground/webhooks.v5 v5.17.0 // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/driver/mysql v1.5.2 // indirect
	gorm.io/gorm v1.25.5 // indirect
	k8s.io/api v0.31.4 // indirect
	k8s.io/apiextensions-apiserver v0.31.2 // indirect
	k8s.io/apiserver v0.31.0 // indirect
	k8s.io/client-go v0.31.4 // indirect
	k8s.io/component-base v0.31.0 // indirect
	k8s.io/component-helpers v0.31.0 // indirect
	k8s.io/klog v1.0.0 // indirect
	k8s.io/klog/v2 v2.130.1 // indirect
	k8s.io/kube-aggregator v0.31.2 // indirect
	k8s.io/kube-openapi v0.0.0-20240228011516-70dd3763d340 // indirect
	k8s.io/kubernetes v1.29.6 // indirect
	k8s.io/metrics v0.29.6 // indirect
	k8s.io/utils v0.0.0-20240711033017-18e509b52bc8 // indirect
	layeh.com/gopher-json v0.0.0-20190114024228-97fed8db8427 // indirect
	moul.io/http2curl v1.0.0 // indirect
	nhooyr.io/websocket v1.8.7 // indirect
	oras.land/oras-go/v2 v2.3.0 // indirect
	sigs.k8s.io/controller-runtime v0.17.2 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/kustomize/api v0.17.2 // indirect
	sigs.k8s.io/kustomize/kustomize/v5 v5.0.4-0.20230601165947-6ce0bf390ce3 // indirect
	sigs.k8s.io/kustomize/kyaml v0.17.1 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.4.4-0.20241211184406-7bf59b3d70ee // indirect
	sigs.k8s.io/yaml v1.4.0 // indirect
)
