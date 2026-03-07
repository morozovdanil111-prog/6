func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Ограничение для парсинга данных формы
	err := r.ParseMultipartForm(10 << 20) // 10MB
	if err != nil {
		log.Printf("Error parsing form: %v", err)
		http.Error(w, "Error processing file", http.StatusInternalServerError)
		return
	}

	// Получаем файл из формы
	file, _, err := r.FormFile("myFile") // имя поля в HTML
	if err != nil {
		log.Printf("Error getting file: %v", err)
		http.Error(w, "Error processing file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Читаем содержимое файла
	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Error reading file: %v", err)
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	// Преобразуем данные с помощью функции AutoConvert
	convertedContent, err := service.AutoConvert(string(fileContent))
	if err != nil {
		log.Printf("Error converting data: %v", err)
		http.Error(w, "Error converting data", http.StatusInternalServerError)
		return
	}

	// Генерация имени файла для записи (с расширением .txt)
	fileName := time.Now().UTC().Format("2006-01-02_15-04-05") + ".txt"

	// Запись в новый файл
	outputFile, err := os.Create(fileName)
	if err != nil {
		log.Printf("Error creating file: %v", err)
		http.Error(w, "Error creating file", http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	// Записываем результат в файл
	_, err = outputFile.WriteString(convertedContent)
	if err != nil {
		log.Printf("Error writing to file: %v", err)
		http.Error(w, "Error writing to file", http.StatusInternalServerError)
		return
	}

	// Возвращаем результат преобразования
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Conversion successful! Result saved to: " + fileName + "\n" + convertedContent))
}
