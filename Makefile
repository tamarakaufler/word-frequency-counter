IMAGE_TAG=v1alpha1
QUAY_PASS?=biggestsecret
WORDCOUNTER_FILE?=./test.txt
WORDCOUNTER_WORKERS?=3

compile:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o wordcounter .

dev: compile
	docker build -t quay.io/tamarakaufler/wordcounter:$(IMAGE_TAG) .

build: dev
	docker login quay.io -u tamarakaufler -p $(QUAY_PASS)
	docker push quay.io/tamarakaufler/wordcounter:$(IMAGE_TAG)

run:
	docker run \
	--name=wordcounter \
	--rm \
	quay.io/tamarakaufler/wordcounter:$(IMAGE_TAG) \
	-file=$(WORDCOUNTER_FILE) \
	-workers=$(WORDCOUNTER_WORKERS)
