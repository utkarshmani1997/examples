NAME     = blog
PID      = /tmp/$(NAME).pid

OUT_DIRS = lib server scripts
BINARIES = $(NAME)-server $(NAME)-cli $(NAME)-tool-db $(NAME)-tool-data

all: gen

build: api
build-all: api tools
gen:
	@geb gen
	@gofmt -w $(OUT_DIRS)
fmt:
	@gofmt -w $(OUT_DIRS)
cloc:
	@echo "----- design -----"
	@cloc design
	@echo "----- subdesigns -----"
	@cloc subdesigns
	@echo "----- all-the-things -----"
	@cloc --not-match-d=vendor $(OUT_DIRS)
.PHONY: all build build-all gen fmt cloc


api:
	@cd server && go build .
	@mv server/server $(NAME)-server
.PHONY: $(NAME)


hot: reload
	@fswatch -r -o -e ".*" -i ".*\.go$$" $(NAME) | xargs -n1 -I{}  make reload || make kill
serve:
	@./$(NAME)
kill:
	@kill `cat $(PID)` || true
reload: kill $(NAME)
	@./$(NAME) & echo $$! > $(PID)
.PHONY: hot serve kill reload


before-tools:
	@echo "Start building tools"
tools: before-tools tool-db tool-data cli
	@echo "Done building tools"
cli:
	@echo "building $(NAME)-cli"
	@cd server/client/$(NAME)-cli && go build .
	@mv server/client/$(NAME)-cli/$(NAME)-cli $(NAME)-cli
tool-db:
	@echo "building $(NAME)-tool-db"
	@cd server/tools/db-ops && go build .
	@mv server/tools/db-ops/db-ops $(NAME)-tool-db
tool-data:
	@echo "building $(NAME)-tool-data"
	@cd server/tools/data-ops && go build .
	@mv server/tools/data-ops/data-ops $(NAME)-tool-data
.PHONY: before-tools tools cli tool-db tool-data

clean: clean-bins clean-output clean-subdesigns
clean-bins:
	@rm -f $(BINARIES)
clean-output:
	@rm -rf *.go config.yaml secrets.yaml $(OUT_DIRS)
clean-subdesigns:
	@rm -rf subdesigns
.PHONY: clean clean-bins clean-output clean-subdesigns


dbs-start:
	@./scripts/db/exec.sh start
dbs-restart:
	@./scripts/db/exec.sh restart
dbs-stop:
	@./scripts/db/exec.sh stop
dbs-wipe:
	@./scripts/db/exec.sh wipe
.PHONY: dbs-start dbs-restart dbs-stop dbs-wipe


app-start:
	@./scripts/app/start.sh
app-stop:
	@./scripts/app/stop.sh
app-restart:
	@./scripts/app/restart.sh
.PHONY: app-start app-stop app-restart

