default: 
	go build -buildmode=c-archive -o ./rbac_c_go.a rbac_c_go.go
	gcc -o a c_go.c rbac_c_go.a -lpthread

clean: 
	rm -f a
	rm -f rbac_c_go.a
	rm -f rbac_c_go.h
