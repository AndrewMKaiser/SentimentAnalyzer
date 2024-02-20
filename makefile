SRC=socialsent.go
EXEC=socialsent
WIN_EXEC=socialsent.exe

all:
	go build $(SRC)

clean:
	rm -f $(EXEC)
	rm -f $(WIN_EXEC)
