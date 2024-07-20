package mapper

import (
	"docto/interfaces"
	"docto/models"
	"strconv"
)

func FileModelToFile(fileModel *models.File) *interfaces.File {
	return &interfaces.File{
		ID:        strconv.FormatInt(int64(fileModel.ID), 10),
		FileName:  fileModel.FileName,
		UpdatedAt: fileModel.UpdatedAt,
		Url:       fileModel.Url,
	}
}

func FileModelsToFiles(fileModels *[]models.File) *[]interfaces.File {
	var files []interfaces.File

	for _, file := range *fileModels {
		files = append(files, *FileModelToFile(&file))
	}

	return &files
}
