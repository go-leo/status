protoc \
		--proto_path=. \
		--proto_path=../leo/ \
		--proto_path=../third_party \
		--status_out=. \
		--status_opt=paths=source_relative \
		*/*.proto