Soporte de análisis Go para CodeQL
Este repositorio de código abierto contiene el extractor, las bibliotecas CodeQL y las consultas que impulsan el soporte de Go en LGTM y los otros productos CodeQL que GitHub pone a disposición de sus clientes en todo el mundo.

Contiene dos componentes principales:

un extractor, escrito en Go, que analiza el código fuente de Go y lo convierte en una base de datos que se puede consultar utilizando CodeQL.
Bibliotecas de análisis estático y consultas escritas en QL que se pueden utilizar para analizar dicha base de datos para encontrar errores de codificación o vulnerabilidades de seguridad.
El objetivo de este proyecto es proporcionar un soporte completo de análisis estático para Go en CodeQL.

Para consultar las consultas y bibliotecas que potencian el soporte de CodeQL para otros lenguajes, visite el repositorio de CodeQL .

Instalación
Simplemente clone este repositorio. No hay dependencias externas.

Si desea utilizar la extensión CodeQL para Visual Studio Code, importe este repositorio en su espacio de trabajo de VS Code.

Uso
Para analizar una base de código de Go, utilice la interfaz de línea de comandos CodeQL para crear una base de datos usted mismo o descargue una base de datos prediseñada de LGTM.com . Luego, puede ejecutar cualquiera de las consultas contenidas en este repositorio, ya sea en la línea de comando o usando la extensión VS Code.

Tenga en cuenta que la rama lgtm.com de este repositorio corresponde a la versión de las consultas que se implementa actualmente en LGTM.com. La rama principal puede contener cambios que aún no se han implementado, por lo que es posible que deba actualizar las bases de datos descargadas de LGTM.com antes de ejecutar consultas sobre ellas.

Contribuciones
¡Las contribuciones son bienvenidas! Consulte nuestras pautas de contribución y nuestro código de conducta para obtener detalles sobre cómo participar en nuestra comunidad.

Licencia
El código de este repositorio tiene la licencia MIT .

Recursos
Escribir consultas CodeQL
Aprendizaje de CodeQL
