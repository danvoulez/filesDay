def generate_text(tipo):
    if tipo == "whitepaper":
        return "Whitepaper gerado automaticamente pelo LLM local."
    elif tipo == "readme":
        return "README gerado automaticamente. Veja instruções para instalação e uso."
    else:
        return f"Texto institucional para {tipo}."

def summarize_logs(logs):
    return f"Sumário automático: {len(logs)} eventos registrados."