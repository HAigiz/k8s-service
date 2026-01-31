# exam

### Настройка подключения к Kubernetes

> Для настройки подключения и работы с кластером Kubernetes необходимы утилиты [task](https://taskfile.dev/docs/installation) и [kubectl](https://kubernetes.io/docs/tasks/tools/). 

Проверить наличие установленных утилит можно следующими командами:
```shell
kubectl version
task --version
```
Для выполнения первоначальной настройки подключения выполните команду:
```shell
task config
```
Проверить подключение можно командой:
```shell
kubectl get all
```
При первом выполнении этой команды ответ сервера должен быть таким:
```
No resources found in <username> namespace.
```

# Итоговая работа


> - Далее в тексте задания \<UserName\> - это ваше имя пользователя на сервере [git.klsh.ru](https://git.klsh.tu)
> - Адрес вашего приложения в сети Internet - http://\<UserName\>.klsh.ru (например [http://matsuev.klsh.ru](http://matsuev.klsh.ru))


В папке ./cmd код четырех приложений на языке Go.
   - root-service
   - api-service
   - info-service
   - login-service

Все приложения принимают на входе значения переменных окружения в виде \*\*\*_APP_NAME и \*\*\*_BIND_PORT (например приложение **root-service** принимает переменные ROOT_APP_NAME и ROOT_BIND_PORT).

### В вашем репозитории на сервере [https://git.klsh.ru](https://git.klsh.ru)

1. Создать папку docker и разместить в ней четыре Dockerfile'a для сборки приложений из исходного кода
2. Для компиляции в контейнерах использовать образ golang:1.24-alpine3.22
3. Для сборки итогового контейнера использовать образ alpine:3.22
4. Готовые контейнеры должны быть опубликованы в реестре git.klsh.ru/\<UserName\>/\<название контейнера\>:latest
5. Сделать сценарии автоматизации для Gitea
   - Lint_and_Test - использовать линтер golangci-lint и команду "go test ./...". Должен запускаться при изменении любых файлов *.go и go.mod
   - Build_Container - компиляция исходного кода и сборка контейнера. Скомпилировать код приложений, упаковать в контейнеры и отправить в репозиторий. Должен запускаться на событие установки тега в виде "v*"
   - Для передачи параметров в сценарии нужно использовать переменные и секреты:
     - vars.REPO_OWNER
     - vars.DOCKER_USERNAME
     - secrets.DOCKER_TOKEN

6. Создать папку deployments и разместить в ней манифесты для Kubernetes

7. Сделать Deployment'ы для Kubernetes (в скобках указано количество реплик приложения):
   - root-app (x2)
   - api-app (x4)
   - info-app (x1)
   - login-app (x1)

8. Обеспечить передачу переменных окружения в приложения через ConfigMap'ы

9. Сделать сервисы для Deployment'ов
   - root-service
   - api-service
   - info-service
   - login-service

10. Сделать Ingress для сервисов
   - "/" -> root-service
   - "/api" -> api-service
   - "/info" -> info-service
   - "/login" -> login-service

> В настройках Ingress в параметре **host** задайте адрес сервера в виде \<UserName\>.klsh.ru. Пример
```yaml
spec:
  rules:
    - host: matsuev.klsh.ru
```

11. Сделать сценарий автоматизации для Gitea
   - Deploy_to_Kubernetes - отправка изменений в кластер Kubernrtes. Должен запускаться на событие "release"
   - Для передачи параметров в сценарий нужно использовать секрет:
     - secrets.KUBE_CONFIG
   - Файл конфигурации для подключения к Kubernetes находится в вашем репозитории и называется kube.config
