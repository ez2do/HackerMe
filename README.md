# Getting started

## How to Build


* In order to successfully build and run your SDK files, you are required to have the following setup in your system:
    * **Go**  (Visit [https://golang.org/doc/install](https://golang.org/doc/install) for more details on how to install Go)
    * **Java VM** Version 8 or later
    * **Eclipse 4.6 (Neon)** or later ([http://www.eclipse.org/neon/](http://www.eclipse.org/neon/))
    * **GoClipse** setup within above installed Eclipse (Follow the instructions at [this link](https://github.com/GoClipse/goclipse/blob/latest/documentation/Installation.md#instructions) to setup GoClipse)
* Ensure that ```GOPATH``` environment variable is set in the system variables. If not, set it to your workspace directory where you will be adding your Go projects.
* The generated code uses unirest-go http library. Therefore, you will need internet access to resolve this dependency. If Go is properly installed and configured, run the following command to pull the dependency:

```Go
go get github.com/apimatic/unirest-go
```

This will install unirest-go in the ```GOPATH``` you specified in the system variables.

Now follow the steps mentioned below to build your SDK:

1. Open eclipse in the Go language perspective and click on the ```Import``` option in ```File``` menu.

![Importing SDK into Eclipse - Step 1](https://apidocs.io/illustration/go?step=import0)

2. Select ```General -> Existing Projects into Workspace``` option from the tree list.

![Importing SDK into Eclipse - Step 2](https://apidocs.io/illustration/go?step=import1)

3. In ```Select root directory```, provide path to the unzipped archive for the generated code. Once the path is set and the Project becomes visible under ```Projects``` click ```Finish```

![Importing SDK into Eclipse - Step 3](https://apidocs.io/illustration/go?step=import2&workspaceFolder=forward%20service-GoLang&projectName=forwardservice_lib)

4. The Go library will be imported and its files will be visible in the Project Explorer

![Importing SDK into Eclipse - Step 4](https://apidocs.io/illustration/go?step=import3&projectName=forwardservice_lib)

## How to Use

The following section explains how to use the ForwardserviceLib library in a new project.

### 1. Add a new Test Project

Create a new project in Eclipse by ```File``` -> ```New``` -> ```Go Project```

![Add a new project in Eclipse](https://apidocs.io/illustration/go?step=createNewProject0)

Name the Project as ```Test``` and click ```Finish```

![Create a new Maven Project - Step 1](https://apidocs.io/illustration/go?step=createNewProject1)

Create a new directory in the ```src``` directory of this project

![Create a new Maven Project - Step 2](https://apidocs.io/illustration/go?step=createNewProject2&projectName=forwardservice_lib)

Name it ```test.com```

![Create a new Maven Project - Step 3](https://apidocs.io/illustration/go?step=createNewProject3&projectName=forwardservice_lib)

Now create a new file inside ```src/test.com```

![Create a new Maven Project - Step 4](https://apidocs.io/illustration/go?step=createNewProject4&projectName=forwardservice_lib)

Name it ```testsdk.go```

![Create a new Maven Project - Step 5](https://apidocs.io/illustration/go?step=createNewProject5&projectName=forwardservice_lib)

In this Go file, you can start adding code to initialize the client library. Sample code to initialize the client library and using its methods is given in the subsequent sections.

### 2. Configure the Test Project

You need to import your generated library in this project in order to make use of its functions. In order to import the library, you can add its path in the ```GOPATH``` for this project. Follow the below steps:

Right click on the project name and click on ```Properties```

![Adding dependency to the client library - Step 1](https://apidocs.io/illustration/go?step=testProject0&projectName=forwardservice_lib)

Choose ```Go Compiler``` from the side menu. Check ```Use project specific settings``` and uncheck ```Use same value as the GOPATH environment variable.```. By default, the GOPATH value from the environment variables will be visible in ```Eclipse GOPATH```. Do not remove this as this points to the unirest dependency.

![Adding dependency to the client library - Step 2](https://apidocs.io/illustration/go?step=testProject1)

Append the library path to this GOPATH

![Adding dependency to the client library - Step 3](https://apidocs.io/illustration/go?step=testProject2&workspaceFolder=forward%20service-GoLang)

Once the path is appended, click on ```OK```

### 3. Build the Test Project

Right click on the project name and click on ```Build Project```

![Build Project](https://apidocs.io/illustration/go?step=buildProject&projectName=forwardservice_lib)

### 4. Run the Test Project

If the build is successful, right click on your Go file and click on ```Run As``` -> ```Go Application```

![Run Project](https://apidocs.io/illustration/go?step=runProject&projectName=forwardservice_lib)

# Class Reference

## <a name="list_of_controllers"></a>List of Controllers

* [shops_pkg](#shops_pkg)
* [products_pkg](#products_pkg)
* [category_pkg](#category_pkg)
* [tiki_pkg](#tiki_pkg)
* [lazada_pkg](#lazada_pkg)

## <a name="shops_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".shops_pkg") shops_pkg

### Get instance

Factory for the ``` SHOPS ``` interface can be accessed from the package shops_pkg.

```go
shops := shops_pkg.NewSHOPS()
```

### <a name="shop_detail"></a>![Method: ](https://apidocs.io/img/method.png ".shops_pkg.ShopDetail") ShopDetail

> TODO: Add a method description


```go
func (me *SHOPS_IMPL) ShopDetail(shopId int64)(,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| shopId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
shopId,_ := strconv.ParseInt("1788160", 10, 8)

var result 
result,_ = shops.ShopDetail(shopId)

```


### <a name="shop_collections"></a>![Method: ](https://apidocs.io/img/method.png ".shops_pkg.ShopCollections") ShopCollections

> TODO: Add a method description


```go
func (me *SHOPS_IMPL) ShopCollections(
            shopId int64,
            from int64)(,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| shopId |  ``` Required ```  | TODO: Add a parameter description |
| from |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
shopId,_ := strconv.ParseInt("1788160", 10, 8)
from,_ := strconv.ParseInt("0", 10, 8)

var result 
result,_ = shops.ShopCollections(shopId, from)

```


### <a name="shop_products"></a>![Method: ](https://apidocs.io/img/method.png ".shops_pkg.ShopProducts") ShopProducts

> TODO: Add a method description


```go
func (me *SHOPS_IMPL) ShopProducts(
            shopId int64,
            from int64)(,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| shopId |  ``` Required ```  | TODO: Add a parameter description |
| from |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
shopId,_ := strconv.ParseInt("1788160", 10, 8)
from,_ := strconv.ParseInt("0", 10, 8)

var result 
result,_ = shops.ShopProducts(shopId, from)

```


### <a name="malls"></a>![Method: ](https://apidocs.io/img/method.png ".shops_pkg.Malls") Malls

> TODO: Add a method description


```go
func (me *SHOPS_IMPL) Malls()(,error)
```

#### Example Usage

```go

var result 
result,_ = shops.Malls()

```


[Back to List of Controllers](#list_of_controllers)

## <a name="products_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".products_pkg") products_pkg

### Get instance

Factory for the ``` PRODUCTS ``` interface can be accessed from the package products_pkg.

```go
products := products_pkg.NewPRODUCTS()
```

### <a name="search_products"></a>![Method: ](https://apidocs.io/img/method.png ".products_pkg.SearchProducts") SearchProducts

> TODO: Add a method description


```go
func (me *PRODUCTS_IMPL) SearchProducts(
            keyword string,
            from int64,
            categoryId int64)(,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| keyword |  ``` Required ```  | TODO: Add a parameter description |
| from |  ``` Required ```  | TODO: Add a parameter description |
| categoryId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
keyword := "may tinh"
from,_ := strconv.ParseInt("0", 10, 8)
categoryId,_ := strconv.ParseInt("13030", 10, 8)

var result 
result,_ = products.SearchProducts(keyword, from, categoryId)

```


### <a name="product_detail"></a>![Method: ](https://apidocs.io/img/method.png ".products_pkg.ProductDetail") ProductDetail

> TODO: Add a method description


```go
func (me *PRODUCTS_IMPL) ProductDetail(
            productId int64,
            shopId int64)(,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| productId |  ``` Required ```  | TODO: Add a parameter description |
| shopId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
productId,_ := strconv.ParseInt("5652805320", 10, 8)
shopId,_ := strconv.ParseInt("239566345", 10, 8)

var result 
result,_ = products.ProductDetail(productId, shopId)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="category_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".category_pkg") category_pkg

### Get instance

Factory for the ``` CATEGORY ``` interface can be accessed from the package category_pkg.

```go
category := category_pkg.NewCATEGORY()
```

### <a name="category_info"></a>![Method: ](https://apidocs.io/img/method.png ".category_pkg.CategoryInfo") CategoryInfo

> TODO: Add a method description


```go
func (me *CATEGORY_IMPL) CategoryInfo(categoryId int64)(,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| categoryId |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
categoryId,_ := strconv.ParseInt("13030", 10, 8)

var result 
result,_ = category.CategoryInfo(categoryId)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="tiki_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".tiki_pkg") tiki_pkg

### Get instance

Factory for the ``` TIKI ``` interface can be accessed from the package tiki_pkg.

```go
tiki := tiki_pkg.NewTIKI()
```

### <a name="get_shop_products"></a>![Method: ](https://apidocs.io/img/method.png ".tiki_pkg.GetShopProducts") GetShopProducts

> TODO: Add a method description


```go
func (me *TIKI_IMPL) GetShopProducts(
            username string,
            page int64)(,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| username |  ``` Required ```  | TODO: Add a parameter description |
| page |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
username := "pham-gia-nguyen"
page,_ := strconv.ParseInt("1", 10, 8)

var result 
result,_ = tiki.GetShopProducts(username, page)

```


[Back to List of Controllers](#list_of_controllers)

## <a name="lazada_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".lazada_pkg") lazada_pkg

### Get instance

Factory for the ``` LAZADA ``` interface can be accessed from the package lazada_pkg.

```go
lazada := lazada_pkg.NewLAZADA()
```

### <a name="m_forward_service_lazada_shop_id_shop_url_https_www_lazada_vn_shop_mo_store1601719537"></a>![Method: ](https://apidocs.io/img/method.png ".lazada_pkg.MForwardServiceLazadaShopIdShopUrlHttpsWwwLazadaVnShopMoStore1601719537") MForwardServiceLazadaShopIdShopUrlHttpsWwwLazadaVnShopMoStore1601719537

> TODO: Add a method description


```go
func (me *LAZADA_IMPL) MForwardServiceLazadaShopIdShopUrlHttpsWwwLazadaVnShopMoStore1601719537(shopUrl string)(,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| shopUrl |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
shopUrl := "https://www.lazada.vn/shop/mo-store-1601719537"

var result 
result,_ = lazada.MForwardServiceLazadaShopIdShopUrlHttpsWwwLazadaVnShopMoStore1601719537(shopUrl)

```


[Back to List of Controllers](#list_of_controllers)



