import os
from flask import Flask, jsonify

app = Flask(__name__)

FILE_PATH = "/home/a.yersinov@corp.onevision.kz/Python_Study/For_Study/Skillbox/Day_4/output_file.txt"  # Фиксированный путь к файлу

def get_summary_rss(file_path):
    units = ["Б", "Кб", "Мб", "Гб", "Тб"]
    
    if not os.path.exists(file_path):
        return {"error": "Файл не найден"}, 404
    
    try:
        with open(file_path, "r") as file:
            lines = file.readlines()[1:]  # Пропускаем заголовок
    except Exception as e:
        return {"error": f"Ошибка чтения файла: {str(e)}"}, 500

    total_rss = 0
    
    for line in lines:
        columns = line.split()
        if len(columns) > 5:  # Проверяем, что строка содержит нужные данные
            try:
                total_rss += int(columns[5])  # RSS находится в 6-м столбце
            except ValueError:
                continue  # Пропускаем строки, если не удаётся преобразовать в число
    
    # Перевод в человекочитаемый формат
    index = 0
    while total_rss >= 1024 and index < len(units) - 1:
        total_rss /= 1024
        index += 1
    
    return {"total_memory": f"{round(total_rss)} {units[index]}"}

@app.route('/processlist', methods=['GET'])
def process_list():
    result = get_summary_rss(FILE_PATH)
    return jsonify(result)

if __name__ == '__main__':
    app.run(port=5000, debug=True)
