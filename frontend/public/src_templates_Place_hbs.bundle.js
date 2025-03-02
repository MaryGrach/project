/*
 * ATTENTION: The "eval" devtool has been used (maybe by default in mode: "development").
 * This devtool is neither made for production nor for readable output files.
 * It uses "eval()" calls to create a separate source file in the browser devtools.
 * If you are trying to read the output file, select a different devtool (https://webpack.js.org/configuration/devtool/)
 * or disable the default devtool with "devtool: false".
 * If you are looking for production-ready output files, see mode: "production" (https://webpack.js.org/configuration/mode/).
 */
(self["webpackChunk"] = self["webpackChunk"] || []).push([["src_templates_Place_hbs"],{

/***/ "./src/templates/Place.hbs":
/*!*********************************!*\
  !*** ./src/templates/Place.hbs ***!
  \*********************************/
/***/ ((module, __unused_webpack_exports, __webpack_require__) => {

eval("var Handlebars = __webpack_require__(/*! ../../node_modules/handlebars/runtime.js */ \"./node_modules/handlebars/runtime.js\");\nfunction __default(obj) { return obj && (obj.__esModule ? obj[\"default\"] : obj); }\nmodule.exports = (Handlebars[\"default\"] || Handlebars).template({\"compiler\":[8,\">= 4.3.0\"],\"main\":function(container,depth0,helpers,partials,data) {\n    var stack1, alias1=container.lambda, alias2=container.escapeExpression, lookupProperty = container.lookupProperty || function(parent, propertyName) {\n        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {\n          return parent[propertyName];\n        }\n        return undefined\n    };\n\n  return \"<div class=\\\"card\\\" id=\\\"card-\"\n    + alias2(alias1(((stack1 = (depth0 != null ? lookupProperty(depth0,\"data\") : depth0)) != null ? lookupProperty(stack1,\"id\") : stack1), depth0))\n    + \"\\\">\\n    <a href=\\\"/sights/\"\n    + alias2(alias1(((stack1 = (depth0 != null ? lookupProperty(depth0,\"data\") : depth0)) != null ? lookupProperty(stack1,\"id\") : stack1), depth0))\n    + \"\\\"><img src=\\\"/\"\n    + alias2(alias1(((stack1 = (depth0 != null ? lookupProperty(depth0,\"data\") : depth0)) != null ? lookupProperty(stack1,\"url\") : stack1), depth0))\n    + \"\\\"></a>\\n    <div class=\\\"card-content\\\">\\n        <p class=\\\"card-name\\\">\"\n    + alias2(alias1(((stack1 = (depth0 != null ? lookupProperty(depth0,\"data\") : depth0)) != null ? lookupProperty(stack1,\"name\") : stack1), depth0))\n    + \"</p>\\n        <p class=\\\"card-description\\\">\"\n    + alias2(alias1(((stack1 = (depth0 != null ? lookupProperty(depth0,\"data\") : depth0)) != null ? lookupProperty(stack1,\"description\") : stack1), depth0))\n    + \"</p>\\n        <div class=\\\"rating\\\">\\n        </div>\\n    </div>\\n</div>\";\n},\"useData\":true});\n\n//# sourceURL=webpack:///./src/templates/Place.hbs?");

/***/ })

}]);