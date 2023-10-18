package client

func MapClientFromClientRepositoryDTO(clientRepositoryDTO ClientRepositoryDTO) Client {
	return Client{
		id:          clientRepositoryDTO.ID,
		description: clientRepositoryDTO.Description,
		password:    clientRepositoryDTO.Password,
	}
}

func MapClientToClientRepositoryDTO(client Client) ClientRepositoryDTO {
	return ClientRepositoryDTO{
		ID:          client.id,
		Description: client.description,
		Password:    client.password,
	}
}
