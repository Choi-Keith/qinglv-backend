SERVICE=User
.PHONY: gen-rpc-ent-logic
gen-rpc-ent-logic: # Generate logic code from Ent, need model and group params | 根据 Ent 生成逻辑代码, 需要设置 model 和 group
	goctls rpc ent --schema=./ent/schema  --style=go_zero --multiple=false --service_name=$(SERVICE) --search_key_num=3 --output=./ --model=$(model) --group=$(group) --proto_out=./pb/$(shell echo $(model) | tr A-Z a-z).proto --overwrite=true
	@echo "Generate logic codes from Ent successfully"