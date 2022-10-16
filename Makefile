IMAGE_SHA=$$(git rev-parse HEAD)
IMAGE_TAG="rg.fr-par.scw.cloud/kiyutink/sowhenthen:${IMAGE_SHA}"

.PHONY: deploy
deploy:
	docker build --tag ${IMAGE_TAG} .
	docker push ${IMAGE_TAG}
	helm upgrade sowhenthen ./helm_chart --set Image.Sha=${IMAGE_SHA} -n sowhenthen --install 
