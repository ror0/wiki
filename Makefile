TARGET		:= wiki
COMMIT		:= $(shell git rev-parse --short HEAD)
GREEN		:= \033[;32m
COLOREND	:= \033[0m

all: $(TARGET)
	@echo -e "$(GREEN)Building project...$(COLOREND)"

$(TARGET):
	@go build -o $(TARGET)

clean:
	@rm -rvf $(TARGET)

.PHONY: all clean
