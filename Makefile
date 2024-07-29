go:
	go run golang/main.go

js:
	npx tsc && npm start

py:
	python -m venv .venv && .venv/Scripts/activate && pip install -r requirements.txt && python -B python/main.py
