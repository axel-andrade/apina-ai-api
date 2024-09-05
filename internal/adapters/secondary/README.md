# Secondary Adapters

## Description (Descrição):

The secondary adapters are responsible for handling interactions with external resources that are not directly related to user input or output, such as databases or third-party services. They encapsulate the logic for accessing and manipulating data from these resources.

(Adaptadores secundários são responsáveis por lidar com interações com recursos externos que não estão diretamente relacionados à entrada ou saída do usuário, como bancos de dados ou serviços de terceiros. Eles encapsulam a lógica para acessar e manipular dados desses recursos.)

## Folder Structure (Estrutura de Pastas):

_database_: Contains the database-related functionality, such as database implementations, mappers, models, and repositories.
(Contém funcionalidades relacionadas a banco de dados, como implementações de banco de dados, mapeadores, modelos e repositórios.)
Reasoning (Justificativa):

Secondary adapters are located in the "secondary" folder because they represent secondary concerns of the application. They are responsible for interacting with external resources that support the application's core functionality but are not directly involved in handling user interactions. Placing them in the "secondary" folder emphasizes their supportive role in providing data storage and access, separate from the primary concerns of user input and output.

(Adaptadores secundários estão localizados na pasta "secondary" porque representam preocupações secundárias da aplicação. Eles são responsáveis por interagir com recursos externos que suportam a funcionalidade principal da aplicação, mas não estão diretamente envolvidos no manuseio de interações do usuário. Colocá-los na pasta "secondary" enfatiza seu papel de suporte em fornecer armazenamento e acesso a dados, separado das preocupações primárias de entrada e saída do usuário.)
