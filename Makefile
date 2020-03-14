## Locations to store google .proto files
google_api_base_location = "vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/"
location = $(google_api_base_location)google/api/
http_out = $(location)http.proto
annotations_out = $(location)annotations.proto

commitid = master
prelink = "https://raw.githubusercontent.com/grpc-ecosystem/grpc-gateway/"
annotation_end = "/third_party/googleapis/google/api/annotations.proto"
http_end = "/third_party/googleapis/google/api/http.proto"
annotation_link = $(prelink)$(commitid)$(annotation_end)
http_link = $(prelink)$(commitid)$(http_end)

.PHONY: proto_cleanup
proto_cleanup: ## Delete generated proto files
	rm -rf proto/gen

.PHONY: proto_fmt
proto_fmt:
	@mkdir -p ${location}
	@wget -q ${annotation_link} -O ${annotations_out}
	@wget -q ${http_link} -O ${http_out}
	prototool format -w proto

.PHONY: proto
proto: proto_fmt ## Generate the proto files
	cd proto/idl/iris-classification && prototool generate

