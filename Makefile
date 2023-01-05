default-data:
	go run ./main.go secrets init --data-dir data-dir
run-macOS:
	./screehavin-macOS-64bit server --data-dir ./data-dir --chain genesis.json --libp2p 0.0.0.0:1478 --nat 192.0.2.1 --seal
run-linux:
	./screehavin-linux-64bit server --data-dir ./data-dir --chain genesis.json --libp2p 0.0.0.0:1478 --nat 192.0.2.1 --seal
run-windows:
	./screehavin-windows-64bit server --data-dir ./data-dir --chain genesis.json --libp2p 0.0.0.0:1478 --nat 192.0.2.1 --seal
batch-staking:
	go run scripts/batch_staking/batch_staking.go