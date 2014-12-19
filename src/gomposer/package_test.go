package gomposer

import (
	"fmt"

	"github.com/icambridge/go-dependency"

	"net/http"
	"net/http/httptest"
	"testing"
	"os"
)

func getMuxAndServer() (*http.ServeMux, *httptest.Server) {

	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	return mux, server
}

func TestPackageRepository_Find(t *testing.T) {
	// todo setUp
	cacheFileName := GetCacheFilename("magetest/magento-behats-extension")
	os.Remove(cacheFileName)
	mux, server := getMuxAndServer()
	apiHit := false
	mux.HandleFunc("/magetest/magento-behats-extension.json", func(w http.ResponseWriter, r *http.Request) {
		if m := "GET"; m != r.Method {
			t.Errorf("Request method = %v, expected %v", r.Method, m)
		}
		fmt.Fprint(w, `{"package":{"name":"magetest\/magento-behat-extension","description":"Magento Behat extension","time":"2012-10-05T19:29:34+00:00","maintainers":[{"name":"md"},{"name":"MarcoDeBortoli"},{"name":"MageTest"},{"name":"alistairstead"}],"versions":{"dev-develop":{"name":"magetest\/magento-behat-extension","description":"Magento Behat extension","keywords":["BDD","Behat","magento"],"homepage":"https:\/\/github.com\/MageTest\/BehatMage","version":"dev-develop","version_normalized":"dev-develop","license":["MIT"],"authors":[{"name":"Alistair Stead","email":"astead@sessiondigital.com"},{"name":"Marcello Duarte","email":"mduarte@sessiondigital.com"},{"name":"Mark Slocock","email":"mark@gpmd.co.uk"},{"name":"Sarunas Valaskevicius","email":"svalaskevicius@sessiondigital.com"},{"name":"Daniel Kidanemariam","email":"dkidanemariam@sessiondigital.com"},{"name":"Other contributors","homepage":"https:\/\/github.com\/MageTest\/BehatMage\/contributors"},{"name":"Marco De Bortoli","email":"marco.debortoli@sessiondigital.com"}],"source":{"type":"git","url":"https:\/\/github.com\/MageTest\/BehatMage.git","reference":"807e7f71509aba930672e1a71487e18963d6cf56"},"dist":{"type":"zip","url":"https:\/\/api.github.com\/repos\/MageTest\/BehatMage\/zipball\/807e7f71509aba930672e1a71487e18963d6cf56","reference":"807e7f71509aba930672e1a71487e18963d6cf56","shasum":""},"type":"library","time":"2014-11-17T20:46:22+00:00","autoload":{"psr-0":{"":["vendor\/magetest\/magento\/src\/app","vendor\/magetest\/magento\/src\/lib","vendor\/magetest\/magento\/src\/app\/code\/local","vendor\/magetest\/magento\/src\/app\/code\/community","vendor\/magetest\/magento\/src\/app\/code\/core"],"MageTest":"src\/"}},"require":{"behat\/mink-extension":"*","behat\/mink-browserkit-driver":"*","behat\/mink-goutte-driver":">=1.0.3","php":"~5.3","behat\/behat":"~2.4.4","sensiolabs\/behat-page-object-extension":"*"},"require-dev":{"mockery\/mockery":"*","magetest\/magento":"*","phpspec\/phpspec":"2.0.*@dev"}},"dev-feature\/Behat3":{"name":"magetest\/magento-behat-extension","description":"Magento Behat extension","keywords":["BDD","Behat","magento"],"homepage":"https:\/\/github.com\/MageTest\/BehatMage","version":"dev-feature\/Behat3","version_normalized":"dev-feature\/Behat3","license":["MIT"],"authors":[{"name":"Alistair Stead","email":"astead@sessiondigital.com"},{"name":"Marcello Duarte","email":"mduarte@sessiondigital.com"},{"name":"Mark Slocock","email":"mark@gpmd.co.uk"},{"name":"Sarunas Valaskevicius","email":"svalaskevicius@sessiondigital.com"},{"name":"Daniel Kidanemariam","email":"dkidanemariam@sessiondigital.com"},{"name":"Other contributors","homepage":"https:\/\/github.com\/MageTest\/BehatMage\/contributors"},{"name":"Marco De Bortoli","email":"marco.debortoli@sessiondigital.com"}],"source":{"type":"git","url":"https:\/\/github.com\/MageTest\/BehatMage.git","reference":"6ff34fe2c43f533617be641ff30f72bf1fed38d5"},"dist":{"type":"zip","url":"https:\/\/api.github.com\/repos\/MageTest\/BehatMage\/zipball\/6ff34fe2c43f533617be641ff30f72bf1fed38d5","reference":"6ff34fe2c43f533617be641ff30f72bf1fed38d5","shasum":""},"type":"library","time":"2014-09-18T16:31:51+00:00","autoload":{"psr-0":{"":["vendor\/magetest\/magento\/src\/app","vendor\/magetest\/magento\/src\/lib","vendor\/magetest\/magento\/src\/app\/code\/local","vendor\/magetest\/magento\/src\/app\/code\/community","vendor\/magetest\/magento\/src\/app\/code\/core"],"MageTest":"src\/"}},"require":{"php":"~5.3","behat\/mink-extension":"*","behat\/mink-browserkit-driver":"*","behat\/mink-goutte-driver":">=1.0.3","sensiolabs\/behat-page-object-extension":"*","behat\/behat":"~3.0"},"require-dev":{"mockery\/mockery":"*","magetest\/magento":"*","phpspec\/phpspec":"~2.0"}},"dev-feature\/Product-image-fixture":{"name":"magetest\/magento-behat-extension","description":"Magento Behat extension","keywords":["BDD","Behat","magento"],"homepage":"https:\/\/github.com\/MageTest\/BehatMage","version":"dev-feature\/Product-image-fixture","version_normalized":"dev-feature\/Product-image-fixture","license":["MIT"],"authors":[{"name":"Alistair Stead","email":"astead@sessiondigital.com"},{"name":"Marcello Duarte","email":"mduarte@sessiondigital.com"},{"name":"Mark Slocock","email":"mark@gpmd.co.uk"},{"name":"Sarunas Valaskevicius","email":"svalaskevicius@sessiondigital.com"},{"name":"Daniel Kidanemariam","email":"dkidanemariam@sessiondigital.com"},{"name":"Other contributors","homepage":"https:\/\/github.com\/MageTest\/BehatMage\/contributors"},{"name":"Marco De Bortoli","email":"marco.debortoli@sessiondigital.com"}],"source":{"type":"git","url":"https:\/\/github.com\/MageTest\/BehatMage.git","reference":"227c020a3ae4df092f53534ed20255a46d7c40e0"},"dist":{"type":"zip","url":"https:\/\/api.github.com\/repos\/MageTest\/BehatMage\/zipball\/227c020a3ae4df092f53534ed20255a46d7c40e0","reference":"227c020a3ae4df092f53534ed20255a46d7c40e0","shasum":""},"type":"library","time":"2014-07-05T19:55:06+00:00","autoload":{"psr-0":{"":["vendor\/magetest\/magento\/src\/app","vendor\/magetest\/magento\/src\/lib","vendor\/magetest\/magento\/src\/app\/code\/local","vendor\/magetest\/magento\/src\/app\/code\/community","vendor\/magetest\/magento\/src\/app\/code\/core"],"MageTest":"src\/"}},"require":{"php":"~5.3","behat\/behat":"~2.4.4","behat\/mink-extension":"*","behat\/mink-browserkit-driver":"*","behat\/mink-goutte-driver":">=1.0.3","sensiolabs\/behat-page-object-extension":"*"},"require-dev":{"phpspec\/phpspec":"2.0.*@dev","mockery\/mockery":"*","magetest\/magento":"*"}},"dev-feature\/spec-magento-context-use-of-fixture-factory":{"name":"magetest\/magento-behat-extension","description":"Magento Behat extension","keywords":["BDD","Behat","magento"],"homepage":"https:\/\/github.com\/MageTest\/BehatMage","version":"dev-feature\/spec-magento-context-use-of-fixture-factory","version_normalized":"dev-feature\/spec-magento-context-use-of-fixture-factory","license":["MIT"],"authors":[{"name":"Alistair Stead","email":"astead@sessiondigital.com"},{"name":"Marcello Duarte","email":"mduarte@sessiondigital.com"},{"name":"Mark Slocock","email":"mark@gpmd.co.uk"},{"name":"Sarunas Valaskevicius","email":"svalaskevicius@sessiondigital.com"},{"name":"Daniel Kidanemariam","email":"dkidanemariam@sessiondigital.com"},{"name":"Other contributors","homepage":"https:\/\/github.com\/MageTest\/BehatMage\/contributors"},{"name":"Marco De Bortoli","email":"marco.debortoli@sessiondigital.com"}],"source":{"type":"git","url":"https:\/\/github.com\/MageTest\/BehatMage.git","reference":"3b2772ff966c898bd3e6eb8d438ccea42ccea9fa"},"dist":{"type":"zip","url":"https:\/\/api.github.com\/repos\/MageTest\/BehatMage\/zipball\/3b2772ff966c898bd3e6eb8d438ccea42ccea9fa","reference":"3b2772ff966c898bd3e6eb8d438ccea42ccea9fa","shasum":""},"type":"library","time":"2014-07-05T09:22:22+00:00","autoload":{"psr-0":{"":["vendor\/magetest\/magento\/src\/app","vendor\/magetest\/magento\/src\/lib","vendor\/magetest\/magento\/src\/app\/code\/local","vendor\/magetest\/magento\/src\/app\/code\/community","vendor\/magetest\/magento\/src\/app\/code\/core"],"MageTest":"src\/"}},"require":{"php":"~5.3","behat\/behat":"~2.4.4","behat\/mink-extension":"*","behat\/mink-browserkit-driver":"*","behat\/mink-goutte-driver":">=1.0.3","sensiolabs\/behat-page-object-extension":"*"},"require-dev":{"phpspec\/phpspec":"2.0.*@dev","mockery\/mockery":"*","magetest\/magento":"*"}},"dev-feature\/exception-thrown-if-sku-is-missing":{"name":"magetest\/magento-behat-extension","description":"Magento Behat extension","keywords":["BDD","Behat","magento"],"homepage":"https:\/\/github.com\/MageTest\/BehatMage","version":"dev-feature\/exception-thrown-if-sku-is-missing","version_normalized":"dev-feature\/exception-thrown-if-sku-is-missing","license":["MIT"],"authors":[{"name":"Alistair Stead","email":"astead@sessiondigital.com"},{"name":"Marcello Duarte","email":"mduarte@sessiondigital.com"},{"name":"Mark Slocock","email":"mark@gpmd.co.uk"},{"name":"Sarunas Valaskevicius","email":"svalaskevicius@sessiondigital.com"},{"name":"Daniel Kidanemariam","email":"dkidanemariam@sessiondigital.com"},{"name":"Other contributors","homepage":"https:\/\/github.com\/MageTest\/BehatMage\/contributors"},{"name":"Marco De Bortoli","email":"marco.debortoli@sessiondigital.com"}],"source":{"type":"git","url":"https:\/\/github.com\/MageTest\/BehatMage.git","reference":"0f6465936f6a5f793b5271c9e0ef2c4d9333b36f"},"dist":{"type":"zip","url":"https:\/\/api.github.com\/repos\/MageTest\/BehatMage\/zipball\/0f6465936f6a5f793b5271c9e0ef2c4d9333b36f","reference":"0f6465936f6a5f793b5271c9e0ef2c4d9333b36f","shasum":""},"type":"library","time":"2014-07-05T06:11:20+00:00","autoload":{"psr-0":{"":["vendor\/magetest\/magento\/src\/app","vendor\/magetest\/magento\/src\/lib","vendor\/magetest\/magento\/src\/app\/code\/local","vendor\/magetest\/magento\/src\/app\/code\/community","vendor\/magetest\/magento\/src\/app\/code\/core"],"MageTest":"src\/"}},"require":{"php":"~5.3","behat\/behat":"~2.4.4","behat\/mink-extension":"*","behat\/mink-browserkit-driver":"*","behat\/mink-goutte-driver":">=1.0.3","sensiolabs\/behat-page-object-extension":"*"},"require-dev":{"phpspec\/phpspec":"2.0.*@dev","mockery\/mockery":"*","magetest\/magento":"*"}},"dev-feature\/created-at-fix":{"name":"magetest\/magento-behat-extension","description":"Magento Behat extension","keywords":["BDD","Behat","magento"],"homepage":"https:\/\/github.com\/MageTest\/BehatMage","version":"dev-feature\/created-at-fix","version_normalized":"dev-feature\/created-at-fix","license":["MIT"],"authors":[{"name":"Alistair Stead","email":"astead@sessiondigital.com"},{"name":"Marcello Duarte","email":"mduarte@sessiondigital.com"},{"name":"Mark Slocock","email":"mark@gpmd.co.uk"},{"name":"Sarunas Valaskevicius","email":"svalaskevicius@sessiondigital.com"},{"name":"Daniel Kidanemariam","email":"dkidanemariam@sessiondigital.com"},{"name":"Other contributors","homepage":"https:\/\/github.com\/MageTest\/BehatMage\/contributors"},{"name":"Marco De Bortoli","email":"marco.debortoli@sessiondigital.com"}],"source":{"type":"git","url":"https:\/\/github.com\/MageTest\/BehatMage.git","reference":"c1f4ce3d96a44bf66c2a804a556102a4ddf9c682"},"dist":{"type":"zip","url":"https:\/\/api.github.com\/repos\/MageTest\/BehatMage\/zipball\/c1f4ce3d96a44bf66c2a804a556102a4ddf9c682","reference":"c1f4ce3d96a44bf66c2a804a556102a4ddf9c682","shasum":""},"type":"library","time":"2014-07-04T18:10:41+00:00","autoload":{"psr-0":{"":["vendor\/magetest\/magento\/src\/app","vendor\/magetest\/magento\/src\/lib","vendor\/magetest\/magento\/src\/app\/code\/local","vendor\/magetest\/magento\/src\/app\/code\/community","vendor\/magetest\/magento\/src\/app\/code\/core"],"MageTest":"src\/"}},"require":{"php":"~5.3","behat\/behat":"~2.4.4","behat\/mink-extension":"*","behat\/mink-browserkit-driver":"*","behat\/mink-goutte-driver":">=1.0.3","sensiolabs\/behat-page-object-extension":"*"},"require-dev":{"phpspec\/phpspec":"2.0.*@dev","mockery\/mockery":"*","magetest\/magento":"*"}},"dev-feature\/exceptions-thrown-from-fixtures":{"name":"magetest\/magento-behat-extension","description":"Magento Behat extension","keywords":["BDD","Behat","magento"],"homepage":"https:\/\/github.com\/MageTest\/BehatMage","version":"dev-feature\/exceptions-thrown-from-fixtures","version_normalized":"dev-feature\/exceptions-thrown-from-fixtures","license":["MIT"],"authors":[{"name":"Alistair Stead","email":"astead@sessiondigital.com"},{"name":"Marcello Duarte","email":"mduarte@sessiondigital.com"},{"name":"Mark Slocock","email":"mark@gpmd.co.uk"},{"name":"Sarunas Valaskevicius","email":"svalaskevicius@sessiondigital.com"},{"name":"Daniel Kidanemariam","email":"dkidanemariam@sessiondigital.com"},{"name":"Other contributors","homepage":"https:\/\/github.com\/MageTest\/BehatMage\/contributors"},{"name":"Marco De Bortoli","email":"marco.debortoli@sessiondigital.com"}],"source":{"type":"git","url":"https:\/\/github.com\/MageTest\/BehatMage.git","reference":"31b5e78abfa6629dc3595e43db912739beaf972b"},"dist":{"type":"zip","url":"https:\/\/api.github.com\/repos\/MageTest\/BehatMage\/zipball\/31b5e78abfa6629dc3595e43db912739beaf972b","reference":"31b5e78abfa6629dc3595e43db912739beaf972b","shasum":""},"type":"library","time":"2014-07-02T10:33:34+00:00","autoload":{"psr-0":{"":["vendor\/magetest\/magento\/src\/app","vendor\/magetest\/magento\/src\/lib","vendor\/magetest\/magento\/src\/app\/code\/local","vendor\/magetest\/magento\/src\/app\/code\/community","vendor\/magetest\/magento\/src\/app\/code\/core"],"MageTest":"src\/"}},"require":{"php":"~5.3","behat\/behat":"~2.4.4","behat\/mink-extension":"*","behat\/mink-browserkit-driver":"*","behat\/mink-goutte-driver":">=1.0.3","sensiolabs\/behat-page-object-extension":"*"},"require-dev":{"phpspec\/phpspec":"2.0.*@dev","mockery\/mockery":"*","magetest\/magento":"*"}},"dev-feature\/fix-date-issue-for-product-fixture":{"name":"magetest\/magento-behat-extension","description":"Magento Behat extension","keywords":["BDD","Behat","magento"],"homepage":"https:\/\/github.com\/MageTest\/BehatMage","version":"dev-feature\/fix-date-issue-for-product-fixture","version_normalized":"dev-feature\/fix-date-issue-for-product-fixture","license":["MIT"],"authors":[{"name":"Alistair Stead","email":"astead@sessiondigital.com"},{"name":"Marcello Duarte","email":"mduarte@sessiondigital.com"},{"name":"Mark Slocock","email":"mark@gpmd.co.uk"},{"name":"Sarunas Valaskevicius","email":"svalaskevicius@sessiondigital.com"},{"name":"Daniel Kidanemariam","email":"dkidanemariam@sessiondigital.com"},{"name":"Other contributors","homepage":"https:\/\/github.com\/MageTest\/BehatMage\/contributors"},{"name":"Marco De Bortoli","email":"marco.debortoli@sessiondigital.com"}],"source":{"type":"git","url":"https:\/\/github.com\/MageTest\/BehatMage.git","reference":"cd0b1861f075c00a7a793d06c916f8be4b225af7"},"dist":{"type":"zip","url":"https:\/\/api.github.com\/repos\/MageTest\/BehatMage\/zipball\/cd0b1861f075c00a7a793d06c916f8be4b225af7","reference":"cd0b1861f075c00a7a793d06c916f8be4b225af7","shasum":""},"type":"library","time":"2014-07-01T22:59:47+00:00","autoload":{"psr-0":{"":["vendor\/magetest\/magento\/src\/app","vendor\/magetest\/magento\/src\/lib","vendor\/magetest\/magento\/src\/app\/code\/local","vendor\/magetest\/magento\/src\/app\/code\/community","vendor\/magetest\/magento\/src\/app\/code\/core"],"MageTest":"src\/"}},"require":{"php":"~5.3","behat\/behat":"~2.4.4","behat\/mink-extension":"*","behat\/mink-browserkit-driver":"*","behat\/mink-goutte-driver":">=1.0.3","sensiolabs\/behat-page-object-extension":"*"},"require-dev":{"phpspec\/phpspec":"2.0.*@dev","mockery\/mockery":"*","magetest\/magento":"*"}},"dev-feature\/refactor":{"name":"magetest\/magento-behat-extension","description":"Magento Behat extension","keywords":["BDD","Behat","magento"],"homepage":"https:\/\/github.com\/MageTest\/BehatMage","version":"dev-feature\/refactor","version_normalized":"dev-feature\/refactor","license":["MIT"],"authors":[{"name":"Alistair Stead","email":"astead@sessiondigital.com"},{"name":"Marcello Duarte","email":"mduarte@sessiondigital.com"},{"name":"Mark Slocock","email":"mark@gpmd.co.uk"},{"name":"Sarunas Valaskevicius","email":"svalaskevicius@sessiondigital.com"},{"name":"Daniel Kidanemariam","email":"dkidanemariam@sessiondigital.com"},{"name":"Other contributors","homepage":"https:\/\/github.com\/MageTest\/BehatMage\/contributors"},{"name":"Marco De Bortoli","email":"marco.debortoli@sessiondigital.com"}],"source":{"type":"git","url":"https:\/\/github.com\/MageTest\/BehatMage.git","reference":"79335a96eab60622341f079f32f481afbcc9b77e"},"dist":{"type":"zip","url":"https:\/\/api.github.com\/repos\/MageTest\/BehatMage\/zipball\/79335a96eab60622341f079f32f481afbcc9b77e","reference":"79335a96eab60622341f079f32f481afbcc9b77e","shasum":""},"type":"library","time":"2013-12-19T12:09:24+00:00","autoload":{"psr-0":{"MageTest":"src\/"}},"require":{"behat\/mink-extension":"*","behat\/mink-browserkit-driver":"*","behat\/mink-goutte-driver":">=1.0.3","sensiolabs\/behat-page-object-extension":"*","php":"~5.3","behat\/behat":"~2.4.4"},"require-dev":{"phpspec\/phpspec":"2.0.*@dev","mockery\/mockery":"*","magetest\/magento":"*"}},"dev-feature\/enhanced-route-excpetion":{"name":"magetest\/magento-behat-extension","description":"Magento Behat extension","keywords":["BDD","Behat","magento"],"homepage":"https:\/\/github.com\/MageTest\/BehatMage","version":"dev-feature\/enhanced-route-excpetion","version_normalized":"dev-feature\/enhanced-route-excpetion","license":["MIT"],"authors":[{"name":"Alistair Stead","email":"astead@sessiondigital.com"},{"name":"Marcello Duarte","email":"mduarte@sessiondigital.com"},{"name":"Mark Slocock","email":"mark@gpmd.co.uk"},{"name":"Sarunas Valaskevicius","email":"svalaskevicius@sessiondigital.com"},{"name":"Daniel Kidanemariam","email":"dkidanemariam@sessiondigital.com"},{"name":"Other contributors","homepage":"https:\/\/github.com\/MageTest\/BehatMage\/contributors"},{"name":"Marco De Bortoli","email":"marco.debortoli@sessiondigital.com"}],"source":{"type":"git","url":"https:\/\/github.com\/MageTest\/BehatMage.git","reference":"f632efee0b8a1948468e84adf1937c301f1b7192"},"dist":{"type":"zip","url":"https:\/\/api.github.com\/repos\/MageTest\/BehatMage\/zipball\/f632efee0b8a1948468e84adf1937c301f1b7192","reference":"f632efee0b8a1948468e84adf1937c301f1b7192","shasum":""},"type":"library","time":"2013-10-09T13:54:52+00:00","autoload":{"psr-0":{"MageTest":"src\/"}},"require":{"php":"~5.3","behat\/behat":"~2.4.4","behat\/mink-extension":"*","behat\/mink-browserkit-driver":"*","behat\/mink-goutte-driver":">=1.0.3"},"require-dev":{"phpspec\/phpspec":"dev-master","mockery\/mockery":"*","magetest\/magento":"*"}},"dev-feature\/travis":{"name":"magetest\/magento-behat-extension","description":"Magento Behat extension","keywords":["BDD","Behat","magento"],"homepage":"https:\/\/github.com\/MageTest\/BehatMage","version":"dev-feature\/travis","version_normalized":"dev-feature\/travis","license":["MIT"],"authors":[{"name":"Alistair Stead","email":"astead@sessiondigital.com"},{"name":"Marcello Duarte","email":"mduarte@sessiondigital.com"},{"name":"Mark Slocock","email":"mark@gpmd.co.uk"},{"name":"Sarunas Valaskevicius","email":"svalaskevicius@sessiondigital.com"},{"name":"Daniel Kidanemariam","email":"dkidanemariam@sessiondigital.com"},{"name":"Other contributors","homepage":"https:\/\/github.com\/MageTest\/BehatMage\/contributors"},{"name":"Marco De Bortoli","email":"marco.debortoli@sessiondigital.com"}],"source":{"type":"git","url":"https:\/\/github.com\/MageTest\/BehatMage.git","reference":"6b29943640f6e504538767f4614204c807144c87"},"dist":{"type":"zip","url":"https:\/\/api.github.com\/repos\/MageTest\/BehatMage\/zipball\/6b29943640f6e504538767f4614204c807144c87","reference":"6b29943640f6e504538767f4614204c807144c87","shasum":""},"type":"library","time":"2013-05-24T16:30:22+00:00","autoload":{"psr-0":{"MageTest":"src\/"}},"require":{"behat\/mink-extension":"*","behat\/mink-browserkit-driver":"*","behat\/mink-goutte-driver":">=1.0.3","php":"~5.3","behat\/behat":"~2.4.4"},"require-dev":{"phpspec\/phpspec":"dev-master","mockery\/mockery":"*","magetest\/magento":"*"}},"dev-cleanup":{"name":"magetest\/magento-behat-extension","description":"Magento Behat extension","keywords":["BDD","Behat","magento"],"homepage":"https:\/\/github.com\/MageTest\/BehatMage","version":"dev-cleanup","version_normalized":"dev-cleanup","license":["MIT"],"authors":[{"name":"Alistair Stead","email":"astead@sessiondigital.com"},{"name":"Marcello Duarte","email":"mduarte@sessiondigital.com"},{"name":"Mark Slocock","email":"mark@gpmd.co.uk"},{"name":"Sarunas Valaskevicius","email":"svalaskevicius@sessiondigital.com"},{"name":"Daniel Kidanemariam","email":"dkidanemariam@sessiondigital.com"},{"name":"Other contributors","homepage":"https:\/\/github.com\/MageTest\/BehatMage\/contributors"}],"source":{"type":"git","url":"https:\/\/github.com\/MageTest\/BehatMage.git","reference":"550b49ff0b8b8defe7f7d31c2829b0b131f16a39"},"dist":{"type":"zip","url":"https:\/\/api.github.com\/repos\/MageTest\/BehatMage\/zipball\/550b49ff0b8b8defe7f7d31c2829b0b131f16a39","reference":"550b49ff0b8b8defe7f7d31c2829b0b131f16a39","shasum":""},"type":"library","time":"2013-05-08T09:26:38+00:00","autoload":{"psr-0":{"MageTest":"src\/","":["vendor\/magento\/app","vendor\/magento\/app\/code\/local","vendor\/magento\/app\/code\/community","vendor\/magento\/app\/code\/core","vendor\/magento\/lib","app","app\/code\/local","app\/code\/community","app\/code\/core","lib"]}},"require":{"behat\/mink-extension":"*","behat\/mink-browserkit-driver":"*","behat\/mink-goutte-driver":">=1.0.3","php":"~5.3","behat\/behat":"~2.4.4"},"require-dev":{"phpspec\/phpspec2":"dev-develop"}},"dev-feature\/licensing":{"name":"magetest\/magento-behat-extension","description":"Magento Behat extension","keywords":["BDD","Behat","magento"],"homepage":"https:\/\/github.com\/MageTest\/BehatMage","version":"dev-feature\/licensing","version_normalized":"dev-feature\/licensing","license":["MIT"],"authors":[{"name":"Alistair Stead","email":"astead@sessiondigital.com"},{"name":"Marcello Duarte","email":"mduarte@sessiondigital.com"},{"name":"Mark Slocock","email":"mark@gpmd.co.uk"},{"name":"Sarunas Valaskevicius","email":"svalaskevicius@sessiondigital.com"},{"name":"Daniel Kidanemariam","email":"dkidanemariam@sessiondigital.com"}],"source":{"type":"git","url":"https:\/\/github.com\/MageTest\/BehatMage.git","reference":"ff2677b3cf161d7ce63b8f45d406fc3d09988049"},"dist":{"type":"zip","url":"https:\/\/api.github.com\/repos\/MageTest\/BehatMage\/zipball\/ff2677b3cf161d7ce63b8f45d406fc3d09988049","reference":"ff2677b3cf161d7ce63b8f45d406fc3d09988049","shasum":""},"type":"library","time":"2013-04-09T18:19:11+00:00","autoload":{"psr-0":{"MageTest":"src\/","":["vendor\/magento\/app","vendor\/magento\/app\/code\/local","vendor\/magento\/app\/code\/community","vendor\/magento\/app\/code\/core","vendor\/magento\/lib","app","app\/code\/local","app\/code\/community","app\/code\/core","lib"]}},"require":{"php":">=5.3.0","behat\/behat":"2.4.4","behat\/mink-extension":"*","behat\/mink-browserkit-driver":"*","behat\/mink-goutte-driver":">=1.0.3"},"require-dev":{"phpspec\/phpspec2":"dev-develop"}},"dev-master":{"name":"magetest\/magento-behat-extension","description":"Magento Behat extension","keywords":["BDD","Behat","magento"],"homepage":"https:\/\/github.com\/MageTest\/BehatMage","version":"dev-master","version_normalized":"9999999-dev","license":["MIT"],"authors":[{"name":"Alistair Stead","email":"astead@sessiondigital.com"},{"name":"Marcello Duarte","email":"mduarte@sessiondigital.com"},{"name":"Mark Slocock","email":"mark@gpmd.co.uk"},{"name":"Sarunas Valaskevicius","email":"svalaskevicius@sessiondigital.com"},{"name":"Daniel Kidanemariam","email":"dkidanemariam@sessiondigital.com"}],"source":{"type":"git","url":"https:\/\/github.com\/MageTest\/BehatMage.git","reference":"370d5efa3ec2472373bf9687ee8b8219a6070dbc"},"dist":{"type":"zip","url":"https:\/\/api.github.com\/repos\/MageTest\/BehatMage\/zipball\/370d5efa3ec2472373bf9687ee8b8219a6070dbc","reference":"370d5efa3ec2472373bf9687ee8b8219a6070dbc","shasum":""},"type":"library","time":"2012-11-08T00:47:14+00:00","autoload":{"psr-0":{"MageTest":"src\/","":["vendor\/magento\/app","vendor\/magento\/app\/code\/local","vendor\/magento\/app\/code\/community","vendor\/magento\/app\/code\/core","vendor\/magento\/lib","app","app\/code\/local","app\/code\/community","app\/code\/core","lib"]}},"require":{"php":">=5.3.0","behat\/behat":"2.4.4","behat\/mink-extension":"*","behat\/mink-browserkit-driver":"*","behat\/mink-goutte-driver":">=1.0.3"},"require-dev":{"phpspec\/phpspec2":"dev-develop"}},"0.0.3":{"name":"magetest\/magento-behat-extension","description":"Magento Behat extension","keywords":["BDD","Behat","magento"],"homepage":"https:\/\/github.com\/MageTest\/BehatMage","version":"0.0.3","version_normalized":"0.0.3.0","license":["MIT"],"authors":[{"name":"Alistair Stead","email":"astead@sessiondigital.com"},{"name":"Marcello Duarte","email":"mduarte@sessiondigital.com"},{"name":"Mark Slocock","email":"mark@gpmd.co.uk"},{"name":"Sarunas Valaskevicius","email":"svalaskevicius@sessiondigital.com"},{"name":"Daniel Kidanemariam","email":"dkidanemariam@sessiondigital.com"}],"source":{"type":"git","url":"https:\/\/github.com\/MageTest\/BehatMage.git","reference":"f2c7865a3c02ff2e106c361848134d86cac438cc"},"dist":{"type":"zip","url":"https:\/\/api.github.com\/repos\/MageTest\/BehatMage\/zipball\/f2c7865a3c02ff2e106c361848134d86cac438cc","reference":"f2c7865a3c02ff2e106c361848134d86cac438cc","shasum":""},"type":"library","time":"2012-10-09T21:48:56+00:00","autoload":{"psr-0":{"MageTest":"src\/","":["vendor\/magento\/lib","vendor\/magento\/app"],"Mage":"vendor\/magento\/app\/code\/core"}},"require":{"php":">=5.3.0","behat\/behat":"2.4.4","behat\/mink-extension":"*","behat\/mink-browserkit-driver":"*","behat\/mink-goutte-driver":"*"},"require-dev":{"phpspec\/phpspec2":"dev-magento-demo"}}},"type":"library","repository":"https:\/\/github.com\/MageTest\/BehatMage","downloads":{"total":4891,"monthly":537,"daily":14},"favers":0}}`)
		apiHit = true
	})

	hc := getHttpClient(server)
	packageRepo := PackageRepository{Client: hc}
	pkg, err := packageRepo.Find("magetest/magento-behats-extension")

	if err != nil {
		t.Errorf("Didn't expect an error but got '%s'", err)
	}

	if apiHit == false {
		t.Errorf("Didn't hit api")
	}
	expectedName := "magetest/magento-behat-extension"
	if pkg.Name != expectedName {
		t.Errorf("Expected '%s' but got '%s'", expectedName, pkg.Name)
	}
	server.Close()
}

func TestPackageRepository_Get(t *testing.T) {

	cacheFileName := GetCacheFilename("m/e")
	os.Remove(cacheFileName)
	mux, server := getMuxAndServer()
	mux.HandleFunc("/m/e.json", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"package":{"name":"m\/e", "versions": {"dev-master": {"name":"m\/e", "version": "dev-master"}, "2.0.0": {"name":"m\/e", "version": "2.0.0"},"2.0.1": {"name":"m\/e", "version": "2.0.1"},"2.1.0": {"name":"m\/e", "version": "2.1.0"},"2.1.1": {"name":"m\/e", "version": "2.1.1"}}}}`)
	})

	hc := getHttpClient(server)
	packageRepo := PackageRepository{Client: hc}
	actual, err := packageRepo.Get("m/e")

	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	exepected := map[string]dependency.Dependency{
		"2.0.0": dependency.Dependency{
			Name:    "m/e",
			Version: "2.0.0",
		},
		"2.0.1": dependency.Dependency{
			Name:    "m/e",
			Version: "2.0.1",
		},
		"2.1.0": dependency.Dependency{
			Name:    "m/e",
			Version: "2.1.0",
		},
		"2.1.1": dependency.Dependency{
			Name:    "m/e",
			Version: "2.1.1",
		},
		"dev-master": dependency.Dependency{
			Name:    "m/e",
			Version: "dev-master",
		},
	}

	for k, v := range exepected {

		if actual[k].Version != v.Version {
			t.Errorf("Expected %q got %q", v, actual[k])
		}
	}
}

func TestPackageRepository_Get_Hits_Api_Once(t *testing.T) {
	cacheFileName := GetCacheFilename("m/e")
	os.Remove(cacheFileName)
	apiHitCount := 0
	mux, server := getMuxAndServer()
	mux.HandleFunc("/m/e.json", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"package":{"name":"m\/e", "versions": {"dev-master": {"name":"m\/e", "version": "dev-master"}, "1.0.0": {"name":"m\/e", "version": "1.0.0"},"1.0.1": {"name":"m\/e", "version": "1.0.1"},"1.1.0": {"name":"m\/e", "version": "1.1.0"},"1.1.1": {"name":"m\/e", "version": "1.1.1"}}}}`)
		apiHitCount++
	})

	hc := getHttpClient(server)
	packageRepo := PackageRepository{Client: hc}
	packageRepo.Get("m/e")
	packageRepo.Get("m/e")

	if expected := 1; apiHitCount != expected {
		t.Errorf("Api was expected to be hit %v times got hit %v", expected, apiHitCount)
	}
}
