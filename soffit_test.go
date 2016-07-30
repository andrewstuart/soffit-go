package soffit

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const request = `{"request":{"windowId":"w_91_ctf4_160.tw","windowState":"exclusive","attributes":{"mode":["view"],"scheme":["https"],"serverName":["portal.astuart.co"],"serverPort":["443"],"secure":["true"]},"parameters":{},"properties":{"content-length":"0","referer":"https://portal.astuart.co/uPortal/p/node-soffit-poc.ctf4/max/render.uP","REMOTE_HOST":"10.255.98.15","REMOTE_ADDR":"10.255.98.15","x-forwarded-port":"443","authorization":"Bearer NjM3MDkzMTktMzAwNi00MTJjLTkzZmUtNmUyNmIzYjFjZDE1","x-forwarded-host":"portal.astuart.co","org.jasig.portal.url.UrlState":"EXCLUSIVE","REQUEST_METHOD":"GET","host":"portal.astuart.co","connection":"close","org.jasig.portal.url.UrlType.RENDER":"true","cache-control":"no-cache","org.jasig.portlet.THEME_NAME":"RespondrJS","x-forwarded-proto":"https","accept-language":"en_US","cookie":"cccMisCode=000; UrlCanonicalizingFilter.REDIRECT_COUNT=1; JSESSIONID=4601E8CDC76754188EA59204D435EC45","org.jasig.portal.url.UrlType":"RENDER","org.jasig.portal.url.UrlState.EXCLUSIVE":"true","x-forwarded-for":"192.168.16.15","pragma":"no-cache","accept":"text/html","x-real-ip":"192.168.16.15","themeName":"RespondrJS","accept-encoding":"gzip, deflate, sdch, br","user-agent":"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.82 Safari/537.36"}},"user":{"username":"rreves@democollege.edu","attributes":{"remoteUser":[],"agentDevice":["Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.82 Safari/537.36"],"eduPersonAffiliation":["student"],"displayName":["Rose Reves"],"givenName":["Rose"],"serverName":["portal.astuart.co"],"impersonating":["false"],"cccMisCodeFromIdp":["000"],"uid":["rreves@democollege.edu"],"cccMisCode":["ZZ1"],"eduPersonPrimaryAffiliation":["student"],"user.login.id":["rreves@democollege.edu"],"eduPersonPrincipalName":["rreves@democollege.edu"],"sn":["Reves"],"username":["rreves@democollege.edu"],"cccId":["ZZZ1111"]},"groups":[]},"context":{"portalInfo":"uPortal/4.3.0-SNAPSHOT","attributes":{},"supportedWindowStates":["detached","normal","maximized","exclusive","minimized"]},"definition":{"title":null,"fname":null,"description":null,"categories":[],"parameters":{},"preferences":{}}}`

func TestUnmarshal(t *testing.T) {
	r := strings.NewReader(request)

	var p Payload

	require.NoError(t, json.NewDecoder(r).Decode(&p))
	require.Equal(t, "w_91_ctf4_160.tw", p.Request.WindowID)
	require.Equal(t, "https", p.Request.Attributes.Get("scheme"))
	require.Equal(t, "https://portal.astuart.co/uPortal/p/node-soffit-poc.ctf4/max/render.uP", p.Request.Properties["referer"])
	require.Equal(t, "Rose Reves", p.User.Attributes.Get("displayName"))
	require.Equal(t, uint64(4), p.Context.PortalInfo.Version.Major)
}
