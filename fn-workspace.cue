module: "namespacelabs.dev/foundation"
requirements: {
	api: 36
}

prebuilts: {
	digest: {
		"namespacelabs.dev/foundation/cmd/fn":                                                         "sha256:f4b9670d2ac29f6a223894bd5b766d9d9f8d30d375e36639375b5734c0a706c3"
		"namespacelabs.dev/foundation/cmd/fncighauth":                                                 "sha256:5d5f242f6a3301b1936a6c91d63ffece58494e5b35795a9214c1bd432b5cb400"
		"namespacelabs.dev/foundation/cmd/fninfo":                                                     "sha256:93854dbc66bc1d02539f547ac284fc5c28be7c165f910d522486ce5df8d89173"
		"namespacelabs.dev/foundation/cmd/fnpipelines":                                                "sha256:1ad3f9e12229a760c91be44dc0b41f0f4e9ef4844369a69e4a8f1140c3038e02"
		"namespacelabs.dev/foundation/devworkflow/web":                                                "sha256:41c069dd9f5bdf634ddb6799202b07655934b6b6e701fc2e0eea97a78fa24a62"
		"namespacelabs.dev/foundation/std/dev/controller":                                             "sha256:36828586f343f79b169bcbfbb026c206bd9886545a60ed765676efbe00848c1d"
		"namespacelabs.dev/foundation/std/development/filesync/controller":                            "sha256:1e7eb5214722e4df2af51b442de658616d6211c1fce0414595b03c4069b90a62"
		"namespacelabs.dev/foundation/std/grpc/httptranscoding/configure":                             "sha256:dad72c2c2f135f6bf4abef844a316e891bef16e2fa09e56afab5671631744ac5"
		"namespacelabs.dev/foundation/std/monitoring/grafana/tool":                                    "sha256:9d1d1b0369906c4eb0ef44ce956f798e3fec8d676d65b705d8d40fcd070b3339"
		"namespacelabs.dev/foundation/std/monitoring/prometheus/tool":                                 "sha256:8a2cbe17314060a7e5029fbe3b9f29cac1bf00abde76289e3e6629f0716e2707"
		"namespacelabs.dev/foundation/std/networking/gateway/controller":                              "sha256:1b8d5df877f8049d6433d9cca64244fc3d718ac0e9f3959611b72c28cc7ddedf"
		"namespacelabs.dev/foundation/std/networking/gateway/server/configure":                        "sha256:3a5dd304b41433cacf33ec999f6084fca2d0fa9e8677610c62078ef9bc17c4d9"
		"namespacelabs.dev/foundation/std/runtime/kubernetes/controller/img":                          "sha256:363c2b61cba8afc4aab2a568ceea1a1bc15130777dec348101109f5d7702888b"
		"namespacelabs.dev/foundation/std/runtime/kubernetes/controller/tool":                         "sha256:826414b9e6861b3133d018e4d722a588fd5287f8f5f793d9cbb7bb267f569ee0"
		"namespacelabs.dev/foundation/std/sdk/buf/baseimg":                                            "sha256:4c411b5f515c9c876d405997e0429325436039a86ac3f431cc5f5554ed925843"
		"namespacelabs.dev/foundation/std/sdk/buf/baseimgnix":                                         "sha256:03e3a1752094b4c9fd314fe5cbf6090d8e5b513096f5456d2fa53425cf8930e9"
		"namespacelabs.dev/foundation/std/secrets/kubernetes":                                         "sha256:143f7665d75dc272c41ac08933fa129ac344a0b2130b43e9102b8c239057d5b0"
		"namespacelabs.dev/foundation/std/startup/testdriver":                                         "sha256:fa5dee56f0f8e22daa5ec275e75cb24f9e65fde36461c6027cb0f72d5b61d41c"
		"namespacelabs.dev/foundation/std/testdata/datastore/keygen":                                  "sha256:985d4f501c9c4b839dd535631a8d104a7fa72c4463d072a8ccec1546530dddd5"
		"namespacelabs.dev/foundation/std/web/http/configure":                                         "sha256:162ffd5eb23b6b69ea51270a7a994e3453f5f9aac9797f9607446e592ce5dedb"
		"namespacelabs.dev/foundation/universe/aws/irsa/prepare":                                      "sha256:219b1c704e11fc6f28221292eab8897c2ea5a3b6cf86dce50090b4a564693d8c"
		"namespacelabs.dev/foundation/universe/aws/s3/internal/configure":                             "sha256:d81aae8a72c420f0fb7732c74cd0da8314636e7b1b068fff7d35e72802bae0b1"
		"namespacelabs.dev/foundation/universe/aws/s3/internal/managebuckets/init":                    "sha256:624bfd42c03bc7d2f6d1152ab0f6226873cd434f01b13e4ea1b6ce0f0793d611"
		"namespacelabs.dev/foundation/universe/db/maria/incluster/tool":                               "sha256:0c03cf9701fbde64405d8025248d945c8326bd322aa8802ce0f0d5e9eef6b632"
		"namespacelabs.dev/foundation/universe/db/maria/init":                                         "sha256:fd01aa9dadf4d27406df2f949714276478fbba626104c1526ed3d76a632fe267"
		"namespacelabs.dev/foundation/universe/db/maria/server/creds/tool":                            "sha256:9bb439b594de2c6969791e56240182015203c9d58cee844a5d158d2e9d7754b1"
		"namespacelabs.dev/foundation/universe/db/postgres/incluster/tool":                            "sha256:74966872e7077a09d827d8f062a67a85b76fd519b7d91497d0b96e64743311e5"
		"namespacelabs.dev/foundation/universe/db/postgres/init":                                      "sha256:ca15983520309cfe9db13ce4f1091ba17353ac790fd9930491babaea7350523c"
		"namespacelabs.dev/foundation/universe/db/postgres/opaque/tool":                               "sha256:6e088995b435dcef1341d15232ca7ec38f8f130ca9d0b36d25e7ea4232f3a913"
		"namespacelabs.dev/foundation/universe/db/postgres/server/creds/tool":                         "sha256:59c189729af5075d5c8c837c4c628751215fcd8c532c85a9511b4d4ccea4cbc6"
		"namespacelabs.dev/foundation/universe/db/postgres/server/img":                                "sha256:9099adfa47466424f0e917f78373ebbcc76714263de7f77722d31bcf230a5ea2"
		"namespacelabs.dev/foundation/universe/development/localstack/s3/internal/configure":          "sha256:5cb1a1cb9b488f86583f2605cc18abb21ff12a9c45b55928d6c1cdcffbbb6ab9"
		"namespacelabs.dev/foundation/universe/development/localstack/s3/internal/managebuckets/init": "sha256:0668e6eba64cfa6463f043804f1e8127e262eaa1f02dce082c082884632cf7a3"
		"namespacelabs.dev/foundation/universe/networking/tailscale/image":                            "sha256:a3d812455cf2ebc50308a0a1886beb2251b4e6ed54ef302705f843e0f8de4020"
		"namespacelabs.dev/foundation/universe/storage/s3/internal/managebuckets":                     "sha256:9c2277a8ae31b0ff60d1fabbee0ead1a6652efae9cf3a61db76d496b6a316021"
		"namespacelabs.dev/foundation/universe/storage/s3/internal/prepare":                           "sha256:91d22af5dfe712874366d972e7714aeeef239611439cb82dfa162c14cd43d749"
	}
	baseRepository: "us-docker.pkg.dev/foundation-344819/prebuilts/"
}
