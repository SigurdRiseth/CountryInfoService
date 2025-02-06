# Country Information Service

Application that provides users with information about countries. The service provides historical population information as a second endpoint.

It uses the following APIs:

* **CountriesNow API** 
  * Endpoint: http://129.241.150.113:3500/api/v0.1/
  * Documentation: https://documenter.getpostman.com/view/1134062/T1LJjU52
* **REST Countries API** 
  * Endpoint: http://129.241.150.113:8080/v3.1/
  * Documentation: http://129.241.150.113:8080/
* Notes:
  * If you just try to explore the endpoints directly, you might not see much (or even see errors). Use the documentation to understand how the services work!
  * Most of those services are publicly available, but please invoke the hosted versions (i.e., the links above) due to load challenges (I rather have my services go down than the public ones :)). Lodge an issue if they do go down (should not be the case, but we never know).

The application is deployed via [Render](https://render.com/) and can be accessed at https://countryrest.onrender.com/.

## Endpoints

Your web service will have three resource root paths:

```
/countryinfo/v1/info/
/countryinfo/v1/population/
/countryinfo/v1/status/
```

Assuming your web service should run on localhost, port 8080, your resource root paths would look something like this:

```
http://localhost:8080/countryinfo/v1/info/
http://localhost:8080/countryinfo/v1/population/
http://localhost:8080/countryinfo/v1/status/
```

A call to any endpoint should display user-readable guidance on how to invoke the service where necessary (not necessary for the status endpoint, of course). However, both setup and use should be documented for in a Readme (use the repository markdown language) alongside your codebase (more details below).

The supported request/response pairs are specified in the following.

For the specifications, the following syntax applies:

* `{:value}` indicates mandatory input parameters specified by the user (i.e., the one using your service).
* `{value}` indicates optional input specified by the user (i.e., the one using your service), where `value' can itself contain further optional input.
* The same notation applies for HTTP parameter specifications (e.g., `{?:param}` is a mandatory parameter, `{?param}` is an optional parameter).

### Country Info Endpoint: Return general country infos

The initial endpoint focuses returns general information for a given country, 2-letter country codes (ISO 3166-2). https://en.wikipedia.org/wiki/ISO_3166-2

#### Request

```
Method: GET
Path: info/{:two_letter_country_code}{?limit=10}
```

* `two_letter_country_code` is the corresponding 2-letter country ISO codes
* `limit` is the number of cities that are listed in the response. The listing of cities should be in ascending alphabetical order. The parameter is optional.

Example request: `info/no`

#### Response

* Content type: `application/json`
* Status code: `200` if everything is OK, appropriate error code otherwise. Ensure to deal with errors gracefully.

Body (Example):

```json
{
	"name": "Norway",
	"continents": ["Europe"],
	"population": 4700000,
	"languages": {"nno":"Norwegian Nynorsk","nob":"Norwegian Bokmål","smi":"Sami"},
	"borders": ["FIN","SWE","RUS"],
	"flag": "https://flagcdn.com/w320/no.png",
	"capital": "Oslo",
	"cities": ["Abelvaer","Adalsbruk","Adland","Agotnes","Agskardet","Aker","Akkarfjord","Akrehamn","Al","Alen","Algard","Almas","Alta","Alvdal","Amli","Amot","Ana-Sira","Andalsnes","Andenes","Angvika","Ankenes","Annstad","Ardal","Ardalstangen","Arendal","Arland","Arneberg","Arnes","Aros","As","Asen","Aseral","Asgardstrand","Ask","Asker","Askim","Aukra","Auli","Aurdal","Aure","Aursmoen","Austbo","Austbygdi","Austevoll","Austmarka","Baerums verk","Bagn","Balestrand","Ballangen","Ballstad","Bangsund","Barkaker","Barstadvik","Batnfjordsora","Batsto","Beisfjord","Beitostolen","Bekkjarvik","Berge","Bergen","Berger","Berkak","Birkeland","Birtavarre","Bjaland","Bjerka","Bjerkvik","Bjoneroa","Bjordal","Bjorke","Bjorkelangen","Bjornevatn","Blaker","Blakset","Bleikvasslia","Bo","Bomlo","Bones","Borge","Borgen","Borhaug","Borkenes","Borregard","Bostad","Bovagen","Boverfjorden","Brandbu","Brandval","Brattholmen","Brattvag","Brekke","Brekstad","Brennasen","Brevik","Bronnoysund","Bru","Bruflat","Brumunddal","Brusand","Bruvik","Bryne","Bud","Burfjord","Buskerud","Buvika","Byglandsfjord","Bygstad","Bykle","Byrknes Nordre","Cavkkus","Dal","Dale","Dalen","Davik","Deknepollen","Digermulen","Dilling","Dimmelsvik","Dirdal","Disena","Dokka","Dolemo","Dovre","Drag","Drammen","Drangedal","Drobak","Dverberg","Dyrvika","Ebru","Egersund","Eggedal","Eggkleiva","Eide","Eidfjord","Eidsa","Eidsberg","Eidsdal","Eidsfoss","Eidsnes","Eidsvag","Eidsvoll","Eidsvoll verk","Eikanger","Eikelandsosen","Eiken","Eina","Eivindvik","Elverum","Enebakkneset","Enga","Engalsvik","Erdal","Erfjord","Ervik","Espeland","Etne","Evanger","Evenskjer","Evje","Eydehavn","Faberg","Faervik","Fagernes","Fagerstrand","Fall","Fardal","Farsund","Fauske","Feda","Fedje","Feiring","Felle","Fenstad","Fetsund","Fevik","Figgjo","Finnoy","Finnsnes","Finsand","Fiska","Fiskum","Fister","Fitjar","Fjellstrand","Fla","Flam","Flateby","Flekke","Flekkefjord","Flemma","Flesberg","Flesnes","Floro","Florvag","Foldereid","Folderoy","Folkestad","Follafoss","Follebu","Follese","Fonnes","Forde","Fornebu","Fosnavag","Fossdalen","Fosser","Fotlandsvag","Fredrikstad","Frekhaug","Fresvik","Frogner","Froland","From","Furnes","Fyrde","Fyresdal","Gan","Gardermoen","Gargan","Garnes","Gasbakken","Gaupen","Geilo","Geithus","Gjerdrum","Gjerstad","Gjolme","Glesvaer","Glomfjord","Godoy","Godvik","Gol","Gran","Gransherad","Granvin","Gratangen","Gravdal","Greaker","Grendi","Gressvik","Grimstad","Groa","Grong","Grua","Gullaug","Gvarv","Haddal","Haegeland","Haerland","Hagan","Hagavik","Hakadal","Halden","Hallingby","Halsa","Haltdalen","Hamar","Hamarvik","Hammerfest","Hansnes","Haram","Hareid","Harstad","Haslum","Hasvik","Hatlestranda","Hauge","Haugesund","Haukeland","Havik","Hebnes","Hedal","Heggedal","Heggenes","Hegra","Heimdal","Helgeland","Helgeroa","Hell","Hellandsjoen","Helleland","Hellesylt","Hellvik","Hemnes","Hemnesberget","Hemnskjela","Hemsedal","Henningsvaer","Herand","Heroysund","Herre","Hersaeter","Hestvika","Hetlevik","Hildre","Hitra","Hjellestad","Hjelmas","Hjelset","Hjorungavag","Hof","Hokkasen","Hokksund","Hol","Hole","Holen","Holmefjord","Holmen","Holmenkollen","Holmestrand","Holsen","Holter","Hommelvik","Hommersak","Honefoss","Hordvik","Hornnes","Horte","Horten","Hov","Hovag","Hovden","Hovet","Hovik verk","Hovin","Hoyanger","Hundven","Hunndalen","Husoy","Hustad","Hvalstad","Hvam","Hvitsten","Hvittingfoss","Hyggen","Hylkje","Hyllestad","Ikornnes","Indre Arna","Indre Billefjord","Indre Klubben","Indre Ulvsvag","Indreby","Innbygda","Inndyr","Innvik","Isdalsto","Ise","Ivgobahta","Jakobselv","Jar","Jaren","Jessheim","Jevnaker","Jomna","Jorpeland","Kabelvag","Kaldfarnes","Kalvag","Kamben","Karasjok","Karlshus","Kaupanger","Kautokeino","Kirkenaer","Kirkenes","Kjeller","Kjellmyra","Kjerstad","Kjollefjord","Kjopsvik","Kleive","Klepp","Kleppe","Kleppesto","Kleppstad","Klofta","Klokkarvik","Knapper","Knappstad","Knarrevik","Knarrlaget","Kolbjornsvik","Kolbotn","Kolbu","Kolltveit","Kolnes","Kolsas","Kolvereid","Kongsberg","Kongshamn","Kongsvika","Kongsvinger","Konsmo","Konsvikosen","Kopervik","Koppang","Korgen","Kornsjo","Korsvegen","Kragero","Krakeroy","Krakstad","Kristiansand","Kristiansund","Kroderen","Krokstadelva","Kval","Kvalsund","Kvam","Kvammen","Kvanne","Kvelde","Kvinesdal","Kvinlog","Kvisvik","Kviteseid","Kyrkjebo","Kyrksaeterora","Lakselv","Laksevag","Laksvatn","Lalm","Land","Langangen","Langesund","Langevag","Langfjordbotn","Langhus","Larkollen","Larvik","Laukvik","Lauvsnes","Lauvstad","Leikang","Leines","Leira","Leirfjord","Leirsund","Leirvik","Leknes","Lena","Lensvik","Lenvik","Lepsoy","Levanger","Lidaladdi","Lier","Lillehammer","Lillesand","Lindas","Loddefjord","Lodingen","Loen","Lofthus","Loken","Lokken Verk","Lom","Lonevag","Longva","Lorenfallet","Loten","Lovund","Lundamo","Lunde","Lunner","Lyngdal","Lyngseidet","Lyngstad","Lysaker","Lysoysundet","Magnor","Malm","Maloy","Malvik","Mandal","Manger","Manndalen","Marheim","Masfjorden","Mathopen","Maura","Mehamn","Meisingset","Melbu","Meldal","Melhus","Melsomvik","Meraker","Mestervik","Midsund","Miland","Minnesund","Mirza Rafi Sauda","Misje","Misvaer","Mjolkeraen","Mjondalen","Mo","Mo i Rana","Modalen","Moelv","Moen","Moi","Molde","Moldjord","Morgedal","Mosby","Mosjoen","Moss","Movik","Myking","Myre","Mysen","Na","Naerbo","Naersnes","Namsos","Namsskogan","Narvik","Naustdal","Nedenes","Nedre Frei","Nesbru","Nesbyen","Nesgrenda","Nesna","Nesoddtangen","Nesttun","Neverdal","Nevlunghamn","Nodeland","Nordby Bruk","Nordfjordeid","Nordkisa","Nordland","Nordstrono","Noresund","Norheimsund","Notodden","Nybergsund","Nyborg","Nydalen","Nygardsjoen","Nyhus","Nykirke","Odda","Odnes","Oksfjord","Oksvoll","Olden","Olderdalen","Olen","Oltedal","Oma","Onarheim","Oppdal","Oppegard","Opphaug","Oresvika","Orje","Orkanger","Ornes","Orre","Os","Oslo","Otta","Otteroy","Ottestad","Oveland","Ovre Ardal","Ovrebo","Oyeren","Oystese","Porsgrunn","Prestfoss","Raholt","Rakkestad","Ramberg","Ramfjordbotn","Ramnes","Rana","Ranasfoss","Randaberg","Ranheim","Raudeberg","Raudsand","Raufoss","Rauland","Re","Reine","Reinsvoll","Reipa","Reistad","Reitan","Rena","Rennebu","Rindal","Ringebu","Ringsaker","Ringstad","Risoyhamn","Rjukan","Roa","Rodberg","Rodoy","Rognan","Rogne","Rokland","Roldal","Rollag","Rolvsoy","Romedal","Rong","Roros","Rorvik","Rosendal","Rossland","Rost","Rovde","Roverud","Royken","Royneberg","Rubbestadneset","Rud","Rygge","Rykene","Rypefjord","Saebo","Saebovik","Saetre","Saevareid","Saeveland","Sagvag","Salhus","Salsbruket","Salsnes","Saltnes","Samuelsberg","Sand","Sandane","Sande","Sandefjord","Sandeid","Sander","Sandnes","Sandnessjoen","Sandshamn","Sandstad","Sandtorg","Sandvika","Sandvoll","Sannidal","Sarpsborg","Saupstad","Selasvatn","Selje","Seljord","Sellebakk","Selva","Selvaer","Sem","Setermoen","Siggerud","Siljan","Silsand","Singsas","Sira","Sirevag","Sistranda","Sjovegan","Skabu","Skage","Skanevik","Skarer","Skarnes","Skatoy","Skaun","Skedsmokorset","Skeie","Ski","Skien","Skjeberg","Skjerstad","Skjervoy","Skjold","Skjoldastraumen","Skjolden","Skodje","Skogn","Skoppum","Skotbu","Skotterud","Skreia","Skudeneshavn","Skulsfjord","Skutvika","Slastad","Slattum","Slemdal","Slemmestad","Sletta","Snaase","Snillfjord","Sogn","Sokna","Sokndal","Soknedal","Sola","Solbergelva","Solvorn","Sommaroy","Somna","Son","Sondeled","Sor-Fron","Sorbo","Soreidgrenda","Sorli","Sortland","Sorum","Sorumsand","Sorvaer","Sorvagen","Sorvik","Spangereid","Sparbu","Sperrebotn","Spillum","Spydeberg","Stabbestad","Stabekk","Stamnes","Stamsund","Stange","Stathelle","Staubo","Stavanger","Stavern","Steigen","Steinberg","Steinkjer","Steinsdalen","Sto","Stokke","Stokmarknes","Stol","Storas","Stordal","Storebo","Storforshei","Storslett","Storsteinnes","Stranda","Straume","Straumen","Strommen","Stronstad","Strusshamn","Stryn","Suldalsosen","Sulisjielmma","Sund","Sundal","Sunde","Sunndalsora","Surnadalsora","Svarstad","Svartskog","Sveio","Svelgen","Svelvik","Svene","Svortland","Sylling","Syvik","Tafjord","Talvik","Tananger","Tanem","Tangen","Tau","Tennevoll","Tennfjord","Tertnes","Tiller","Tingvoll","Tistedal","Tjeldsto","Tjelta","Tjong","Tjorvag","Tjotta","Tofte","Tolga","Tomasjorda","Tomter","Tonstad","Tornes","Torod","Torp","Torpo","Tovik","Trana","Tranby","Trengereid","Tretten","Treungen","Trofors","Trollfjorden","Tromsdalen","Trondheim","Trones","Turoy","Tvedestrand","Tveit","Tynset","Tyristrand","Tysnes","Tysse","Tyssedal","Uggdal","Ulefoss","Ulstein","Ulsteinvik","Ulvagen","Ulvik","Undeim","Uskedalen","Utsira","Utskarpen","Uvdal","Vadheim","Vage","Vagland","Vaksdal","Vale","Valen","Valer","Valestrand","Valestrandfossen","Valldal","Valle","Valsoyfjord","Vangsvika","Vannvag","Vanse","Varangerbotn","Varhaug","Vassenden","Vatne","Vedavagen","Vegarshei","Veggli","Venabygd","Vennesla","Verdal","Vestby","Vestfossen","Vestnes","Vestra Mosterhamn","Vestre Gausdal","Vevang","Vevelstad","Vigrestad","Vikebygd","Vikedal","Vikersund","Vikesa","Vikran","Vingelen","Vinje","Vinstra","Voksa","Volda","Vollen","Vormedal","Vormsund","Voss","Vossestrand","Vraliosen","Ytre Alvik","Olavtoppen","Kapp Valdivia","Kapp Circoncision","Nyrøysa","Kapp Norvegia","Larsøya","Kapp Fie","Cape Lollo","Thompson Island","Åneby","Årnes","Ås","Aurskog-Høland","Bærum","Billingstad","Bjørkelangen","Blakstad","Drøbak","Enebakk","Fet","Fjellfoten","Frogn","Hurdal","Kløfta","Lillestrøm","Lørenskog","Nannestad","Nes","Neskollen","Nesodden","Nittedal","Oppegård","Råholt","Rælingen","Rotnes","Skedsmo","Skui","Sørum","Sørumsand","Ullensaker","Ål","Åros","Flå","Hønefoss","Hurum","Krødsherad","Modum","Nedre Eiker","Nore og Uvdal","Øvre Eiker","Ringerike","Røyken","Sætre","Sigdal","Skoger","Ávanuorri","Båtsfjord","Berlevåg","Bjørnevatn","Gamvik","Honningsvåg","Kárášjohka","Kjøllefjord","Lebesby","Loppa","Måsøy","Nesseby","Nordkapp","Øksfjord","Porsanger","Sør-Varanger","Tana","Vadsø","Vardø","Åmot","Åsnes","Eidskog","Engerdal","Folldal","Grue","Kirkenær","Løten","Nord-Odal","Rendalen","Sør-Odal","Spetalen","Stor-Elvdal","Trysil","Våler","Ågotnes","Askøy","Austrheim","Bømlo","Fjell","Fusa","Jondal","Kinsarvik","Knappskog","Knarvik","Kvinnherad","Lindås","Lonevåg","Meland","Mosterhamn","Osterøy","Øygarden","Øystese","Radøy","Sagvåg","Samnanger","Sandsli","Skogsvågen","Stord","Storebø","Syfteland","Ullensvang","Ytre Arna","Ytrebygda","Ålesund","Åndalsnes","Averøy","Batnfjordsøra","Brattvåg","Eidsvåg","Elnesvågen","Fræna","Giske","Gjemnes","Herøy","Hopen","Larsnes","Nesset","Norddal","Nordstranda","Ørskog","Ørsta","Rauma","Rensvik","Sandøy","Sjøholt","Smøla","Steinshamn","Sula","Sunndal","Sunndalsøra","Surnadal","Sykkylven","Tomra","Ulsteinvik weather pws station","Vanylven","Alstahaug","Andøy","Beiarn","Bindal","Bodø","Bogen","Bø","Brønnøy","Brønnøysund","Dønna","Evenes","Evjen","Flakstad","Gildeskål","Gladstad","Grane","Hadsel","Hamarøy","Hattfjelldal","Hauknes","Kabelvåg","Kjøpsvik","Leland","Løding","Lødingen","Løpsmarka","Lurøy","Meløy","Mosjøen","Moskenes","Øksnes","Ørnes","Rødøy","Røst","Saltdal","Sandnessjøen","Sømna","Sørfold","Sørland","Svolvær","Terråk","Tjeldsund","Træna","Tysfjord","Vågan","Værøy","Vefsn","Vega","Vestvågøy","Vik","Dombås","Etnedal","Fossbergom","Gausdal","Gjøvik","Hundorp","Lesja","Nord-Aurdal","Nord-Fron","Nordre Land","Østre Toten","Øyer","Øystre Slidre","Sel","Skjåk","Søndre Land","Sør-Aurdal","Sør-Fron","Vågå","Vågåmo","Vang","Vestre Slidre","Vestre Toten","Sjølyststranda","Aremark","Fossby","Hobøl","Hvaler","Lervik","Marker","Ørje","Råde","Rømskog","Ryggebyen","Skiptvet","Skjærhalden","Trøgstad","Åkrehamn","Bjerkreim","Bokn","Eigersund","Eike","Finnøy","Forsand","Gjesdal","Hå","Hauge i Dalane","Hjelmeland","Hommersåk","Jørpeland","Judaberg","Karmøy","Kvitsøy","Lund","Lyefjell","Nærbø","Ølen","Rennesøy","Sauda","Sæveland","Strand","Suldal","Time","Tysvær","Vedavågen","Vikeså","Vikevåg","Vindafjord","Årdal","Årdalstangen","Askvoll","Aurland","Bremanger","Eid","Farnes","Fjaler","Flora","Florø","Førde","Gaular","Gaupne","Gloppen","Gulen","Hardbakke","Hermansverk","Hornindal","Høyanger","Jølster","Lærdal","Lærdalsøyri","Leikanger","Luster","Måløy","Sogndal","Solund","Vågsøy","Bamble","Hjartdal","Kragerø","Nissedal","Nome","Prestestranda","Sauherad","Tinn","Tokke","Balsfjord","Bardu","Berg","Dyrøy","Gryllefjord","Ibestad","Kåfjord","Karlsøy","Kvæfjord","Kvænangen","Lavangen","Lyngen","Målselv","Nordreisa","Salangen","Sjøvegan","Skånland","Skjervøy","Sørreisa","Storfjord","Torsken","Tranøy","Tromsø","Å i Åfjord","Åfjord","Agdenes","Berkåk","Bjugn","Botngård","Fillan","Flatanger","Fosnes","Frosta","Frøya","Hemne","Holtålen","Høylandet","Inderøy","Indre Fosen","Klæbu","Kyrksæterøra","Leka","Lierne","Meråker","Midtre Gauldal","Namdalseid","Nærøy","Orkdal","Osen","Overhalla","Ørland","Raarvihke - Røyrvik","Ranemsletta","Roan","Røros","Rørvik","Røyrvik","Selbu","Snåase","Snåase - Snåsa","Stjørdal","Stjørdalshalsen","Tydal","Verran","Vikna","Åseral","Audnedal","Hægebostad","Justvik","Liknes","Lindesnes","Marnardal","Sirdal","Skålevik","Songdalen","Søgne","Strai","Vestbygd","Vigeland","Årøysund","Åsgårdstrand","Barkåker","Færder","Gullhaug","Selvik","Tjøme","Tønsberg"]
}
```

* **Note**: Returning all cities can lead to considerable delay in response (potentially a large number of values!). For test cases, it is best to contrain this initially to ensure reasonable response times.
* **Hint**: Think about some known test cases first ('known' as in 'you know the results'), before you develop. This way you know the results your service should provide, which reduces the opportunity for bugs in your codebase. Also beware of some oddities that you may only discover during developing or testing the service (a typical challenge of dealing with real-world APIs!).

### Country Population Endpoint: Return population levels for given time frames

The second endpoint should return population levels for individual years for a given country (identified based on country code), as well as the mean value of those. Optionally, the endpoint should allow you to limit the number of returned values by time frames. Otherwise, all values are returned.

#### Request

```
Method: GET
Path: population/{:two_letter_country_code}{?limit={:startYear-endYear}}
```

* `{:two_letter_country_code}` refers to the ISO 3166-2 identifier of the country.
* `{?limit={:startYear-endYear}}` is an optional parameter that constrains the population history to values between start year and end year (boundary values are included).

Example requests:

* `population/no`
* `population/no?limit=2000-2010`

#### Response

* Content type: `application/json`
* Status code: `200` if everything is OK, appropriate error code otherwise. Ensure to deal with errors gracefully.

Body (Example):
```json
{
   "mean": 5044396,
   "values": [
	        {"year":2010,"value":4889252},
	        {"year":2011,"value":4953088},
	        {"year":2012,"value":5018573},
	        {"year":2013,"value":5079623},
	        {"year":2014,"value":5137232},
	        {"year":2015,"value":5188607}
             ]
}
```

* **Note**: The mean value is rounded to full integer, the values are the individual values for the corresponding years.

### Diagnostics Endpoint: Getting a status overview of services

The diagnostics interface indicates the availability of individual services this service depends on. The reporting occurs based on status codes returned by the dependent services, and it further provides information about the uptime of the service.

#### Request

```
Method: GET
Path: status/
```

#### Response

* Content type: `application/json`
* Status code: `200` if everything is OK, appropriate error code otherwise.

Body:
```json
{
   "countriesnowapi": "<http status code for CountriesNow API>",
   "restcountriesapi": "<http status code for RestCountries API>",
   "version": "v1",
   "uptime": <time in seconds since the last re/start of your service>
}
```

* **Note**: `<some value> indicates placeholders for values to be populated by the service. The values for the different services can either be string (with a content generated based on status code) or integer (with the value being the status code value itself). For this endpoint, the spec is a bit more generous, since it should mostly serve you as a developer.

## Deployment

The service is to be deployed on [Render](https://render.com/). You will need to provide the URL to the deployed service as part of the submission. You will also need to ensure that your service is running after the submission deadline (before that, you may suspend the service, so as to avoid exceeding the execution limits for free Render accounts - we will discuss this in class).

* Hint: Start the development locally first and deploy once you are confident about individual or all endpoints.

## General Aspects

As indicated during the initial sessions, ensure you work with professionalism in mind (see Course Rules). In addition to professionalism, you are at liberty to introduce further features into your service, as long it does not break the specification given above.

Please work in the provided workspace environment (see here - create an issue if you have trouble accessing it) for your user and create a project `assignment-1` in this workspace (so it is easiest to find if we look for it).

Consider to review the example projects provided as part of the lectures and coding tutorials in order to develop understanding of concepts, rather (or in addition) to online resources. Chances are that you will have a better basic understanding, before you consult internet resources (e.g., StackOverflow) for more specialised questions.

As mentioned above, be sensitive to rate limits of external services. If needed, consider _mocking_ the remote services during development.

Where possible, avoid the use of third-party libraries for this assignment. The functionality of this assignment can be developed using the Golang standard API without any problem. We have discussed the reasons for this in class.

## Submission

The assignment is an individual assignment. The submission deadline is provided on the course main wiki page. Extensions to the deadline are handled according to the Course Rules. (Please acquaint yourself with those early!)

As part of the submission you will need to provide:

* a link to your code repository (ensure it is set to the visibility `internal` at that stage to enable peer review)
* a link to the deployed Render service

In addition, we will provide you with an option to clarify aspects of your submission (e.g., aspects that don't quite work, or additional features).

The submission occurs via our submission system that not only facilitates the submission, but also the peer review of the assignment. Instructions for the submission system (submission, review) will be introduced in class, and provided here at a later stage.

## Peer Review

After the submission deadline, there will be a second deadline during which you will review other students' submissions. To do this, the system provides you with a checklist of aspects to assess. You will need to provide sincere reviews for at least two submissions to meet the mandatory requirements of peer review, but you can review as many submissions as you like, which counts towards your participation mark for the course (quality trumps quantity!). The peer-review deadline will be indicated closer to submission time and then listed on the main course wiki page.