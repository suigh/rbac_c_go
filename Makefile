default: 
	go build -buildmode=c-archive -o ./rbac_c_go.a rbac_c_go.go

clean: 
	rm -f rbac_c_go.a
	rm -f rbac_c_go.h
