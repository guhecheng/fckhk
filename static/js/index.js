$(function () {
    $("#create_code").on("click", function () {
        let content = $("input[name='input_content']").val();
        if (content == '') {
            return false;
        }

        $.ajax({
            type: 'get',
            dataType: 'json',
            url : '/qrCode/getQrcode',
            data: {'qrcode_content' : content},
            success: function (response) {
                if (response['code'] == 0) {
                    $("#code_image").attr('src', response['qrCode_url']);
                    $("#download_image").attr('download', 'http://127.0.0.1:8080' + response['qrCode_url']);
                }
            }
        });
    });
    $("#get_content").on('click', function () {
        let content = $("input[name='content']").val();
        if (content == '') {
            return false;
        }

        $.ajax({
            type: 'get',
            dataType: 'json',
            url : '/qrCode/getCode',
            data: {'content' : content},
            success: function (response) {
                if (response['code'] == 0) {
                    let image = "<img src='" + response['qrCode_url'] + "' />"
                    $("#image_content").empty().append(image)
                }
            }
        });
    });

    $("#input_js").on("blur", function () {
        let input_js = $(this).val();
        if ($.trim(input_js) == '') {
            return false;
        }


    });
    // 工具方法
    var formatJson = function(json, options) {
        var reg = null,
            formatted = '',
            pad = 0,
            PADDING = '    '; // one can also use '\t' or a different number of spaces
        // optional settings
        options = options || {};
        // remove newline where '{' or '[' follows ':'
        options.newlineAfterColonIfBeforeBraceOrBracket = (options.newlineAfterColonIfBeforeBraceOrBracket === true) ? true : false;
        // use a space after a colon
        options.spaceAfterColon = (options.spaceAfterColon === false) ? false : true;

        // begin formatting...

        // make sure we start with the JSON as a string
        if (typeof json !== 'string') {
            json = JSON.stringify(json);
        }
        // parse and stringify in order to remove extra whitespace
        json = JSON.parse(json);
        json = JSON.stringify(json);

        // add newline before and after curly braces
        reg = /([\{\}])/g;
        json = json.replace(reg, '\r\n$1\r\n');

        // add newline before and after square brackets
        reg = /([\[\]])/g;
        json = json.replace(reg, '\r\n$1\r\n');

        // add newline after comma
        reg = /(\,)/g;
        json = json.replace(reg, '$1\r\n');

        // remove multiple newlines
        reg = /(\r\n\r\n)/g;
        json = json.replace(reg, '\r\n');

        // remove newlines before commas
        reg = /\r\n\,/g;
        json = json.replace(reg, ',');

        // optional formatting...
        if (!options.newlineAfterColonIfBeforeBraceOrBracket) {
            reg = /\:\r\n\{/g;
            json = json.replace(reg, ':{');
            reg = /\:\r\n\[/g;
            json = json.replace(reg, ':[');
        }
        if (options.spaceAfterColon) {
            reg = /\:/g;
            json = json.replace(reg, ': ');
        }

        $.each(json.split('\r\n'), function (index, node) {
            var i = 0,
                indent = 0,
                padding = '';

            if (node.match(/\{$/) || node.match(/\[$/)) {
                indent = 1;
            } else if (node.match(/\}/) || node.match(/\]/)) {
                if (pad !== 0) {
                    pad -= 1;
                }
            } else {
                indent = 0;
            }

            for (i = 0; i < pad; i++) {
                padding += PADDING;
            }
            formatted += padding + node + '\r\n';
            pad += indent;
        });
        return formatted;
    }

});