# data-platform-api-plant-creates-rmq-kube

data-platform-api-plant-creates-rmq-kube は、周辺業務システム　を データ連携基盤 と統合することを目的に、API でプラントデータを登録するマイクロサービスです。  
https://xxx.xxx.io/api/API_PLANT_SRV/creates/

## 動作環境

data-platform-api-plant-creates-rmq-kube の動作環境は、次の通りです。  
・ OS: LinuxOS （必須）  
・ CPU: ARM/AMD/Intel（いずれか必須）  


## 本レポジトリ が 対応する API サービス
data-platform-api-plant-creates-rmq-kube が対応する APIサービス は、次のものです。

APIサービス URL: https://xxx.xxx.io/api/API_PLANT_SRV/creates/

## 本レポジトリ に 含まれる API名
data-platform-api-plant-creates-rmq-kube には、次の API をコールするためのリソースが含まれています。  

* A_General（プラント - 基本データ）
* A_StorageLocation（プラント - 在庫保管場所データ）

## API への 値入力条件 の 初期値
data-platform-api-plant-creates-rmq-kube において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

## データ連携基盤のAPIの選択的コール

Latona および AION の データ連携基盤 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"General" が指定されています。    
  
```
	"api_schema": "DPFMPlantCreates",
	"accepter": ["General"],
	"plant": null,
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "DPFMPlantCreates",
	"accepter": ["All"],
	"plant": null,
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて DPFM_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *DPFMAPICaller) AsyncPlantCreates(
	accepter []string,
	input *dpfm_api_input_reader.SDC,

	log *logger.Logger,

) []error {
	wg := sync.WaitGroup{}
	mtx := sync.Mutex{}
	errs := make([]error, 0, 5)

	sqlUpdateFin := make(chan error)

	for _, fn := range accepter {
		wg.Add(1)
		switch fn {
		case "General":
			go c.General(&wg, &mtx, sqlUpdateFin, log, &errs, input)
		case "StorageLocation":
			go c.StorageLocation(&wg, &mtx, sqlUpdateFin, log, &errs, input)
		default:
			wg.Done()
		}
	}

	ticker := time.NewTicker(10 * time.Second)
	select {
	case e := <-sqlUpdateFin:
		if e != nil {
			mtx.Lock()
			errs = append(errs, e)
			return errs
		}
	case <-ticker.C:
		mtx.Lock()
		errs = append(errs, xerrors.New("time out"))
		return errs
	}

	return nil
}
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-data-platform](https://github.com/latonaio/golang-logging-library-for-data-platform) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は プラント の 基本データ が取得された結果の JSON の例です。  
以下の項目のうち、"BusinessPartner"〜"IsMarkedForDeletion" は、/DPFM_API_Output_Formatter/type.go 内 の Type General {} による出力結果です。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。  

```
{
	"connection_key": "request",
	"result": true,
	"redis_key": "abcdefg",
	"filepath": "/var/lib/aion/Data/rededge_sdc/abcdef.json",
	"api_status_code": 200,
	"runtime_session_id": "boi9ar543dg91ipdnspi099u231280ab0v8af0ew",
	"business_partner": 201,
	"service_label": "PLANT",
	"Plant": {
		"BusinessPartner": 101,
		"Plant": "SS99",
		"PlantFullName": "",
		"PlantName": "ABC社",
		"Language": "JA",
		"CreationDate": "2022-10-20",
		"CreationTime": "2022-10-20T12:10:16",
		"LastChangeDate": "2022-11-20",
		"LastChangeTime": "2022-11-20T14:11:15",
		"PlantFoundationDate": "",
		"PlantLiquidationDate": "",
		"SearchTerm1": "",
		"SearchTerm2": "",
		"PlantDeathDate": "",
		"PlantIsBlocked": "",
		"GroupPlantName1": "",
		"GroupPlantName2": "",
		"AddressID": null,
		"Country": "",
		"TimeZone": "",
		"PlantIDByExtSystem": "",
		"IsMarkedForDeletion": null
	},
	"api_schema": "DPFMPlantCreates",
	"accepter": ["Plant"],
	"order_id": null,
	"deleted": false,
	"sql_update_result": true,
	"sql_update_error": "",
	"subfunc_result": null,
	"subfunc_error": "",
	"exconf_result": null,
	"exconf_error": "",
	"api_processing_result": true,
	"api_processing_error": ""
}
```
