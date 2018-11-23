/* bootstrap */
(function bootstrap() {
    return JSON.stringify(eval(func)(JSON.parse(json)));
})();

/**
 *
 * @param j
 * @returns {*}
 */
function onStartScenario(j) {
    j.log = 0;
    return j;
}

/**
 *
 * @param j
 * @returns {*}
 */
function onEndScenario(j) {
    j.log = "aaaaa";
    throw "Error Error Error";
    return j;
}

/**
 *
 * @param j
 * @returns {*}
 */
function setRequestParameter(j) {
    return j;
}

/**
 *
 * @param j
 * @returns {*}
 */
function setRequestHeader(j) {
    return j;
}

/**
 *
 * @param j
 * @returns {*}
 */
function saveParams(j) {
    return j;
}
