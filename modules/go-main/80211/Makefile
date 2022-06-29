include /usr/local/etc/PcapPlusPlus.mk

# All Target
all:
	g++ $(PCAPPP_BUILD_FLAGS) $(PCAPPP_INCLUDES) -c -o main.o main.cpp
	g++ $(PCAPPP_LIBS_DIR) -static-libstdc++ -o main main.o $(PCAPPP_LIBS)
	g++ $(PCAPPP_BUILD_FLAGS) $(PCAPPP_INCLUDES) -c -o interface.o interfaces.cpp
	g++ $(PCAPPP_LIBS_DIR) -static-libstdc++ -o interface interface.o $(PCAPPP_LIBS)
	go build mod/goul/goul.go

# Clean Target
clean:
	rm main.o
	rm main
