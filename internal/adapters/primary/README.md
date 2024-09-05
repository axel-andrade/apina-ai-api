# Primary Adapters

## Description (Descrição):

The primary adapters are responsible for handling interactions with the external world, such as handling HTTP requests, presenting data to clients, and defining routes. They primarily deal with the application's input and output.

(Adaptadores primários são responsáveis por lidar com interações com o mundo externo, como manipular requisições HTTP, apresentar dados aos clientes e definir rotas. Eles lidam principalmente com a entrada e saída da aplicação.)

## Folder Structure (Estrutura de Pastas):

_http_: Contains the HTTP-related functionality, such as controllers, middlewares, presenters, routes, and the HTTP server configuration.
(Contém funcionalidades relacionadas ao HTTP, como controladores, middlewares, apresentadores, rotas e a configuração do servidor HTTP.)
schemas: Stores JSON schemas used for request validation and documentation purposes.
(Armazena esquemas JSON usados para validação de requisições e documentação.)
README.md: Documentation or instructions specific to the primary adapters.
(Documentação ou instruções específicas para os adaptadores primários.)
Reasoning (Justificativa):

Primary adapters are located in the "primary" folder because they represent the primary entry points of the application. They interact directly with external systems or clients, such as web browsers or mobile applications. These adapters encapsulate the mechanisms for receiving requests, processing them, and sending responses back to the clients. Placing them in the "primary" folder emphasizes their role as the interface between the application and its users or external systems.

(Adaptadores primários estão localizados na pasta "primary" porque representam os principais pontos de entrada da aplicação. Eles interagem diretamente com sistemas externos ou clientes, como navegadores da web ou aplicativos móveis. Esses adaptadores encapsulam os mecanismos para receber requisições, processá-las e enviar respostas de volta aos clientes. Colocá-los na pasta "primary" enfatiza seu papel como interface entre a aplicação e seus usuários ou sistemas externos.)
