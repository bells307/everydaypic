---
title: everydaypic API v
language_tabs:
  - "": ""
language_clients:
  - "": ""
toc_footers: []
includes: []
search: false
highlight_theme: darkula
headingLevel: 2

---

<!-- Generator: Widdershins v4.0.1 -->

<h1 id="">everydaypic API v</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

<h1 id="-image">image</h1>

## Добавить изображение

<a id="opIdcreate-image"></a>

> Code samples

`POST /v1/image`

> Body parameter

```yaml
file: string
fileName: string
name: string

```

<h3 id="добавить-изображение-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|object|true|none|
|» file|body|string(binary)|true|Бинарные данные файла|
|» fileName|body|string|true|Имя файла|
|» name|body|string|true|Имя изображения|

> Example responses

> 200 Response

```json
[
  {
    "created": "string",
    "fileName": "string",
    "id": "string",
    "name": "string",
    "userID": "string"
  }
]
```

<h3 id="добавить-изображение-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Изображение успешно создано|Inline|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Неправильно сформирован запрос|string|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Внутренняя ошибка сервиса|string|

<h3 id="добавить-изображение-responseschema">Response Schema</h3>

Status Code **200**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|[[model.Image](#schemamodel.image)]|false|none|none|
|» created|string|false|none|Дата создания|
|» fileName|string|false|none|Имя файла|
|» id|string|false|none|ID изображения|
|» name|string|false|none|Имя изображения|
|» userID|string|false|none|ID пользователя, добавившего изображение|

<aside class="success">
This operation does not require authentication
</aside>

## Получить информация об изображениях

<a id="opIdget-images-info"></a>

> Code samples

`GET /v1/image/info`

<h3 id="получить-информация-об-изображениях-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|fileName|query|array[string]|false|Имя файла|
|id|query|array[string]|false|ID изображения|

> Example responses

> 200 Response

```json
[
  {
    "created": "string",
    "fileName": "string",
    "id": "string",
    "name": "string",
    "userID": "string"
  }
]
```

<h3 id="получить-информация-об-изображениях-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Информация об изображении|Inline|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Неправильно сформирован запрос|string|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|Изображение не найдено|string|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Внутренняя ошибка сервиса|string|

<h3 id="получить-информация-об-изображениях-responseschema">Response Schema</h3>

Status Code **200**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|[[model.Image](#schemamodel.image)]|false|none|none|
|» created|string|false|none|Дата создания|
|» fileName|string|false|none|Имя файла|
|» id|string|false|none|ID изображения|
|» name|string|false|none|Имя изображения|
|» userID|string|false|none|ID пользователя, добавившего изображение|

<aside class="success">
This operation does not require authentication
</aside>

## Скачать изображение

<a id="opIddownload-image"></a>

> Code samples

`GET /v1/image/{id}`

<h3 id="скачать-изображение-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|id|path|string|true|ID изображения|

> Example responses

> 200 Response

<h3 id="скачать-изображение-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Бинарные данные изображения|string|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|Изображение не найдено|string|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Внутренняя ошибка сервиса|string|

<aside class="success">
This operation does not require authentication
</aside>

# Schemas

<h2 id="tocS_model.Image">model.Image</h2>
<!-- backwards compatibility -->
<a id="schemamodel.image"></a>
<a id="schema_model.Image"></a>
<a id="tocSmodel.image"></a>
<a id="tocsmodel.image"></a>

```json
{
  "created": "string",
  "fileName": "string",
  "id": "string",
  "name": "string",
  "userID": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|created|string|false|none|Дата создания|
|fileName|string|false|none|Имя файла|
|id|string|false|none|ID изображения|
|name|string|false|none|Имя изображения|
|userID|string|false|none|ID пользователя, добавившего изображение|

