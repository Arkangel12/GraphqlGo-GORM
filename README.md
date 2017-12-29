# GraphqlGo-GORM


La idea principal de este repo es para ayudar tanto a mi como al resto 
de la comunidad en Go a entender como configurar Graphql. <br /> 

Sobre la carpeta de config se cuenta con un archivo json el cual ayudara con
la configuracion de la BD que se utilizara ya que se utiliza GORM son libres 
de hacer uso de a BD de su preferecia. <br /> 

<i>Importante: cambiar nombre de config.example.json a config.json </i>

El package de mail solo se quedo por que era partede la pruebas pueden hacer 
caso omiso de el ya que no es necesario. 

Este es un primero paso, solo comprende lo basico para arrancar la base de datos
no cuenta con un cliente de graphql como <b> Apollo </b> ni nada similar pero si 
con una interfaz para explorar la base de datos e interactuar con ella.

Si no cuentan con una base de datos yo estoy utilizando Postgres para este ejemplo
pueden tener una DB de prueba como yo creando una cuenta <a href="https://www.elephantsql.com/">ElephanSQL</a>
les regalara una TinyDB de hasta 20 MB que debe ser suficiente para el ejemplo.

Continuara...