define build_binaries
	@echo "- Building binaries..."
	@GOOS=linux GOARCH=amd64 go build -o bin/fetchAllExpenses lambdas/fetchAllExpenses/main.go
	@echo "Finished building binaries"
endef

define zip_files
	@echo "- Zipping files..."
	@for file in bin/*; do \
		mv $$file ./bin/bootstrap; \
		zip -j $$file.zip ./bin/bootstrap; \
		rm ./bin/bootstrap; \
	done
	@echo "Finished zipping files"
endef

define clean_up
	@echo "- Cleaning up..."
	@rm -rf bin
endef

define deploy_to_aws
	@echo "- Deploying to AWS..."
	@serverless deploy --stage dev
	@echo "Finished deploying to AWS"
endef

deploy:
	@rm -rf bin/
	${build_binaries}
	${zip_files}
	${deploy_to_aws}
	
