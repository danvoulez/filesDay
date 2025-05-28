import argparse
import sys
import json
from datetime import datetime

def logline(event_type, payload):
    log_entry = {
        "timestamp": datetime.utcnow().isoformat() + "Z",
        "system": "loglineos",
        "event": event_type,
        "payload": payload
    }
    print(json.dumps(log_entry))

def run_benchmark():
    # Simulação de benchmark
    result = {"script": "bench_insert", "result": "ok", "latency_ms": 1.2}
    logline("benchmark_exec", result)

def run_poc():
    # Simulação de prova de conceito pagamento
    result = {"tipo": "pix", "valor": 100, "status": "confirmado"}
    logline("poc_pagamento", result)

def main():
    parser = argparse.ArgumentParser(description="LogLineOS CLI")
    parser.add_argument("command", choices=["benchmark", "poc"], help="Comando a executar")
    args = parser.parse_args()
    
    if args.command == "benchmark":
        run_benchmark()
    elif args.command == "poc":
        run_poc()
    else:
        parser.print_help()
        sys.exit(1)

if __name__ == "__main__":
    main()