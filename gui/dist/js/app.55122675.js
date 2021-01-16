(function(e){function t(t){for(var n,s,u=t[0],c=t[1],i=t[2],d=0,p=[];d<u.length;d++)s=u[d],Object.prototype.hasOwnProperty.call(o,s)&&o[s]&&p.push(o[s][0]),o[s]=0;for(n in c)Object.prototype.hasOwnProperty.call(c,n)&&(e[n]=c[n]);l&&l(t);while(p.length)p.shift()();return a.push.apply(a,i||[]),r()}function r(){for(var e,t=0;t<a.length;t++){for(var r=a[t],n=!0,u=1;u<r.length;u++){var c=r[u];0!==o[c]&&(n=!1)}n&&(a.splice(t--,1),e=s(s.s=r[0]))}return e}var n={},o={app:0},a=[];function s(t){if(n[t])return n[t].exports;var r=n[t]={i:t,l:!1,exports:{}};return e[t].call(r.exports,r,r.exports,s),r.l=!0,r.exports}s.m=e,s.c=n,s.d=function(e,t,r){s.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:r})},s.r=function(e){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},s.t=function(e,t){if(1&t&&(e=s(e)),8&t)return e;if(4&t&&"object"===typeof e&&e&&e.__esModule)return e;var r=Object.create(null);if(s.r(r),Object.defineProperty(r,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var n in e)s.d(r,n,function(t){return e[t]}.bind(null,n));return r},s.n=function(e){var t=e&&e.__esModule?function(){return e["default"]}:function(){return e};return s.d(t,"a",t),t},s.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},s.p="/";var u=window["webpackJsonp"]=window["webpackJsonp"]||[],c=u.push.bind(u);u.push=t,u=u.slice();for(var i=0;i<u.length;i++)t(u[i]);var l=c;a.push([0,"chunk-vendors"]),r()})({0:function(e,t,r){e.exports=r("56d7")},"034f":function(e,t,r){"use strict";r("85ec")},"56d7":function(e,t,r){"use strict";r.r(t);r("e260"),r("e6cf"),r("cca6"),r("a79d");var n=r("2b0e"),o=function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("div",{attrs:{id:"app"}},[r("h1",[e._v("Last api calls")]),r("table",[e._m(0),e._l(e.users,(function(t){return r("tr",{key:t.id},[r("td",[e._v(e._s(t.id))]),r("td",[e._v(e._s(t.email))]),r("td",[e._v(e._s(t.token))]),r("td",[e._v(e._s(t.requests))])])}))],2),r("hr"),r("table",[e._m(1),e._l(e.calls,(function(t){return r("tr",{key:t.id},[r("td",[e._v(e._s(t.id))]),r("td",[r("a",{attrs:{href:t.url,title:t.url,target:"_blank"}},[e._v("link")])]),r("td",[e._v(e._s(t.user.email))])])}))],2),r("hr"),r("h1",[e._v("Proxy CRUD")]),r("select",{directives:[{name:"model",rawName:"v-model",value:e.form.type,expression:"form.type"}],on:{change:function(t){var r=Array.prototype.filter.call(t.target.options,(function(e){return e.selected})).map((function(e){var t="_value"in e?e._value:e.value;return t}));e.$set(e.form,"type",t.target.multiple?r:r[0])}}},[r("option",{attrs:{value:"http"}},[e._v("http")]),r("option",{attrs:{value:"sock4"}},[e._v("sock4")]),r("option",{attrs:{value:"sock5"}},[e._v("sock5")])]),r("input",{directives:[{name:"model",rawName:"v-model",value:e.form.address,expression:"form.address"}],attrs:{size:"40",type:"text",placeholder:"login:password@address:port"},domProps:{value:e.form.address},on:{input:function(t){t.target.composing||e.$set(e.form,"address",t.target.value)}}}),r("input",{directives:[{name:"model",rawName:"v-model",value:e.form.country,expression:"form.country"}],attrs:{size:"20",type:"text",placeholder:"ru,ua,de,us,uk etc",maxlength:"3"},domProps:{value:e.form.country},on:{keypress:function(t){return!t.type.indexOf("key")&&e._k(t.keyCode,"enter",13,t.key,"Enter")?null:e.sendProxy(t)},input:function(t){t.target.composing||e.$set(e.form,"country",t.target.value)}}}),r("button",{on:{click:e.sendProxy}},[e._v("submit")]),r("table",[e._m(2),e._l(e.proxies,(function(t){return r("tr",{key:t.id},[r("td",[e._v(e._s(t.id))]),r("td",[e._v(e._s(t.type))]),r("td",[e._v(e._s(t.address))]),r("td",[e._v(e._s(t.port))]),r("td",[e._v(e._s(t.login))]),r("td",[e._v(e._s(t.password))]),r("td",[e._v(e._s(t.country))]),r("td",{staticClass:"remove",on:{click:function(r){return e.rmProxy(t.id)}}},[e._v("x")])])}))],2)])},a=[function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("tr",[r("th",[e._v("user.id")]),r("th",[e._v("user.email")]),r("th",[e._v("user.token")]),r("th",[e._v("user.requests")])])},function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("tr",[r("th",[e._v("id")]),r("th",[e._v("url")]),r("th",[e._v("user")])])},function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("tr",[r("th",[e._v("id")]),r("th",[e._v("type")]),r("th",{staticStyle:{"min-width":"250px"}},[e._v("address")]),r("th",[e._v("port")]),r("th",[e._v("login")]),r("th",[e._v("password")]),r("th",[e._v("country")]),r("th",[e._v("delete")])])}],s=(r("d3b7"),r("ac1f"),r("1276"),r("96cf"),r("1da1")),u={name:"App",data:function(){return{auth:!1,proxies:null,calls:null,users:null,form:{type:"http",address:null,country:null}}},methods:{sendProxy:function(){var e=this;return Object(s["a"])(regeneratorRuntime.mark((function t(){var r,n,o,a,s,u,c,i,l,d;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return r=Object.assign({},e.$data.form),n=r.address.split("@"),o=n[0].split(":"),a=n[1].split(":"),s=o[0],u=o[1],c=a[0],i=a[1],console.log(e.$data.form.country,e.$data.form.type,s,u,c,i),t.next=11,fetch("http://localhost:8090/proxy/create",{method:"POST",credentials:"include",headers:{Accept:"application/json","Content-Type":"application/json"},body:JSON.stringify({country:e.$data.form.country,type:e.$data.form.type,login:s,password:u,address:c,port:i})});case 11:return l=t.sent,t.next=14,l.json();case 14:return d=t.sent,"ok"!==d.status&&alert("some error: "+d.msg),t.next=18,e.updateProxies();case 18:case"end":return t.stop()}}),t)})))()},rmProxy:function(e){var t=this;return Object(s["a"])(regeneratorRuntime.mark((function r(){var n,o;return regeneratorRuntime.wrap((function(r){while(1)switch(r.prev=r.next){case 0:return r.next=2,fetch("http://localhost:8090/proxy/".concat(e,"/delete"),{method:"GET",credentials:"include"});case 2:return n=r.sent,r.next=5,n.json();case 5:return o=r.sent,"ok"!==o.status&&alert("some error"),r.next=9,t.updateProxies();case 9:case"end":return r.stop()}}),r)})))()},updateProxies:function(){var e=this;return Object(s["a"])(regeneratorRuntime.mark((function t(){var r;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,fetch("http://localhost:8090/proxies",{method:"GET",credentials:"include"});case 2:if(r=t.sent,200!==r.statusCode){t.next=6;break}return e.$data.auth=!0,t.abrupt("return");case 6:return t.next=8,r.json();case 8:e.$data.proxies=t.sent;case 9:case"end":return t.stop()}}),t)})))()},getCalls:function(){var e=this;return Object(s["a"])(regeneratorRuntime.mark((function t(){var r,n;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,fetch("http://localhost:8090/calls",{method:"GET",credentials:"include"});case 2:return r=t.sent,t.next=5,r.json();case 5:return e.$data.calls=t.sent,t.next=8,fetch("http://localhost:8090/users",{method:"GET",credentials:"include"});case 8:return n=t.sent,t.next=11,n.json();case 11:e.$data.users=t.sent;case 12:case"end":return t.stop()}}),t)})))()}},mounted:function(){var e=this;return Object(s["a"])(regeneratorRuntime.mark((function t(){return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,e.updateProxies();case 2:return t.next=4,e.getCalls();case 4:case"end":return t.stop()}}),t)})))()}},c=u,i=(r("034f"),r("2877")),l=Object(i["a"])(c,o,a,!1,null,null,null),d=l.exports;n["a"].config.productionTip=!1,new n["a"]({render:function(e){return e(d)}}).$mount("#app")},"85ec":function(e,t,r){}});
//# sourceMappingURL=app.55122675.js.map