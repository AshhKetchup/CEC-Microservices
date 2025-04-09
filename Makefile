# Root Makefile
.PHONY: build-all test-all clean-all

SERVICES := auth product order delivery gateway

build-all:
	@for service in $(SERVICES); do \
		$(MAKE) -C services/$$service build; \
	done

test-all:
	@for service in $(SERVICES); do \
		$(MAKE) -C services/$$service test; \
	done

clean-all:
	@for service in $(SERVICES); do \
		$(MAKE) -C services/$$service clean; \
	done

docker-build-all:
	@for service in $(SERVICES); do \
		docker-compose build $$service-service; \
	done

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down