/*
 * ATTENTION: The "eval" devtool has been used (maybe by default in mode: "development").
 * This devtool is neither made for production nor for readable output files.
 * It uses "eval()" calls to create a separate source file in the browser devtools.
 * If you are trying to read the output file, select a different devtool (https://webpack.js.org/configuration/devtool/)
 * or disable the default devtool with "devtool: false".
 * If you are looking for production-ready output files, see mode: "production" (https://webpack.js.org/configuration/mode/).
 */
(self["webpackChunk"] = self["webpackChunk"] || []).push([["src_templates_Link_hbs"],{

/***/ "./src/templates/Link.hbs":
/*!********************************!*\
  !*** ./src/templates/Link.hbs ***!
  \********************************/
/***/ ((module, __unused_webpack_exports, __webpack_require__) => {

eval("var Handlebars = __webpack_require__(/*! ../../node_modules/handlebars/runtime.js */ \"./node_modules/handlebars/runtime.js\");\nfunction __default(obj) { return obj && (obj.__esModule ? obj[\"default\"] : obj); }\nmodule.exports = (Handlebars[\"default\"] || Handlebars).template({\"1\":function(container,depth0,helpers,partials,data) {\n    var lookupProperty = container.lookupProperty || function(parent, propertyName) {\n        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {\n          return parent[propertyName];\n        }\n        return undefined\n    };\n\n  return \"        <img src=\\\"\"\n    + container.escapeExpression(container.lambda((depth0 != null ? lookupProperty(depth0,\"src\") : depth0), depth0))\n    + \"\\\">\\n\";\n},\"compiler\":[8,\">= 4.3.0\"],\"main\":function(container,depth0,helpers,partials,data) {\n    var stack1, alias1=container.lambda, alias2=container.escapeExpression, lookupProperty = container.lookupProperty || function(parent, propertyName) {\n        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {\n          return parent[propertyName];\n        }\n        return undefined\n    };\n\n  return \" <div class=\\\"\"\n    + alias2(alias1((depth0 != null ? lookupProperty(depth0,\"className\") : depth0), depth0))\n    + \"\\\">\\n\"\n    + ((stack1 = lookupProperty(helpers,\"if\").call(depth0 != null ? depth0 : (container.nullContext || {}),(depth0 != null ? lookupProperty(depth0,\"src\") : depth0),{\"name\":\"if\",\"hash\":{},\"fn\":container.program(1, data, 0),\"inverse\":container.noop,\"data\":data,\"loc\":{\"start\":{\"line\":2,\"column\":4},\"end\":{\"line\":4,\"column\":11}}})) != null ? stack1 : \"\")\n    + \"    <a id=\\\"\"\n    + alias2(alias1((depth0 != null ? lookupProperty(depth0,\"id\") : depth0), depth0))\n    + \"\\\" href=\\\"\"\n    + alias2(alias1((depth0 != null ? lookupProperty(depth0,\"url\") : depth0), depth0))\n    + \"\\\">\\n        \"\n    + alias2(alias1((depth0 != null ? lookupProperty(depth0,\"label\") : depth0), depth0))\n    + \"\\n    </a>\\n</div>\";\n},\"useData\":true});\n\n//# sourceURL=webpack:///./src/templates/Link.hbs?");

/***/ })

}]);