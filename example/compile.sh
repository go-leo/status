protoc \
		--proto_path=. \
		--proto_path=../leo/ \
		--proto_path=../third_party \
		--go_out=. \
		--go_opt=paths=source_relative \
		--status_out=. \
		--status_opt=paths=source_relative \
		*/*.proto