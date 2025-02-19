"use strict";(self.webpackChunkgrafana_lokiexplore_app=self.webpackChunkgrafana_lokiexplore_app||[]).push([[599],{4599:(e,t,n)=>{n.r(t),n.d(t,{default:()=>O,updatePlugin:()=>y});var r=n(8531),a=n(7781),i=n(1269),l=n(6089),o=n(2007),c=n(5959),s=n.n(c),u=n(3241),p=n(2871);function d(e,t,n,r,a,i,l){try{var o=e[i](l),c=o.value}catch(e){return void n(e)}o.done?t(c):Promise.resolve(c).then(r,a)}function m(e){return function(){var t=this,n=arguments;return new Promise((function(r,a){var i=e.apply(t,n);function l(e){d(i,r,a,l,o,"next",e)}function o(e){d(i,r,a,l,o,"throw",e)}l(void 0)}))}}function v(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}const g=e=>({colorWeak:l.css`
    color: ${e.colors.text.secondary};
  `,marginTop:l.css`
    margin-top: ${e.spacing(3)};
  `,marginTopXl:l.css`
    margin-top: ${e.spacing(6)};
  `,label:(0,l.css)({display:"flex",alignItems:"center",marginBottom:e.spacing(.75)}),icon:(0,l.css)({marginLeft:e.spacing(1)})}),f=function(){var e=m((function*(e,t){try{yield y(e,t),r.locationService.reload()}catch(e){p.v.error(e,{msg:"Error while updating the plugin"})}}));return function(t,n){return e.apply(this,arguments)}}(),b={container:"data-testid ac-container",interval:"data-testid ac-interval-input",submit:"data-testid ac-submit-form"},y=function(){var e=m((function*(e,t){const n=(0,r.getBackendSrv)().fetch({url:`/api/plugins/${e}/settings`,method:"POST",data:t});return(yield(0,i.lastValueFrom)(n)).data}));return function(t,n){return e.apply(this,arguments)}}(),h=e=>{try{if(e){const t=a.rangeUtil.intervalToSeconds(e);return(0,u.isNumber)(t)&&t>=3600}return!0}catch(e){}return!1},O=({plugin:e})=>{const t=(0,o.useStyles2)(g),{enabled:n,pinned:r,jsonData:a}=e.meta;var i,l;const[u,p]=(0,c.useState)({interval:null!==(i=null==a?void 0:a.interval)&&void 0!==i?i:"",isValid:h(null!==(l=null==a?void 0:a.interval)&&void 0!==l?l:"")});return s().createElement("div",{"data-testid":b.container},s().createElement(o.FieldSet,{label:"Settings"},s().createElement(o.Field,{invalid:!h(u.interval),error:'Interval is invalid. Please enter an interval longer then "60m". For example: 3d, 1w, 1m',description:s().createElement("span",null,"The maximum interval that can be selected in the time picker within the Explore Logs app. If empty, users can select any time range interval in Explore Logs. ",s().createElement("br",null),"Example values: 7d, 24h, 2w"),label:"Maximum time picker interval",className:t.marginTop},s().createElement(o.Input,{width:60,id:"interval","data-testid":b.interval,label:"Max interval",value:null==u?void 0:u.interval,placeholder:"7d",onChange:e=>{const t=e.target.value.trim();var n,r;p((n=function(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{},r=Object.keys(n);"function"==typeof Object.getOwnPropertySymbols&&(r=r.concat(Object.getOwnPropertySymbols(n).filter((function(e){return Object.getOwnPropertyDescriptor(n,e).enumerable})))),r.forEach((function(t){v(e,t,n[t])}))}return e}({},u),r=null!=(r={interval:t,isValid:h(t)})?r:{},Object.getOwnPropertyDescriptors?Object.defineProperties(n,Object.getOwnPropertyDescriptors(r)):function(e){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t.push.apply(t,n)}return t}(Object(r)).forEach((function(e){Object.defineProperty(n,e,Object.getOwnPropertyDescriptor(r,e))})),n))}})),s().createElement("div",{className:t.marginTop},s().createElement(o.Button,{type:"submit","data-testid":b.submit,onClick:()=>f(e.meta.id,{enabled:n,pinned:r,jsonData:{interval:u.interval}}),disabled:!h(u.interval)},"Save settings"))))}}}]);
//# sourceMappingURL=599.js.map