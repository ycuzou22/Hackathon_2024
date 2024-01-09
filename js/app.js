$(function() {

    if(document.cookie.match('(^|;) ?\_acceptcookies=([^;]*)(;|$)')==null) {
        $('.cookiebar').fadeIn();
        $('#cookiebar_close').click(function(){
            $('.cookiebar').fadeOut();
            var expire = new Date();
            expire.setYear(expire.getFullYear()+1);
            document.cookie='_acceptcookies=true; expires='+expire.toUTCString()+'; path=/';
        });
    }
});

$(function () {
    /** data-confirm on links **/
    $('a[data-confirm="true"]').click(function (event) {
        if (!confirm(_trans.appConfirm)) {
            event.preventDefault();
        }
    });

    /** bootstrap tooltips **/
    $('[data-toggle="tooltip"]').tooltip();

    $('#filter-form input').change(function () {
        $('#filter-form').submit();
    });

    var fixHelperModified = function (e, tr) {
        var $originals = tr.children();
        var $helper = tr.clone();
        $helper.children().each(function (index) {
            $(this).width($originals.eq(index).width());
            $(this).css('background', '#eee').css('border-bottom', '1px solid #ddd')
                                             .css('opacity','0.8');
        });
        return $helper;
    },
        updateIndex = function (e, ui) {
            $('td.index', ui.item.parent()).each(function (i) {
                var ordre = i++;
                $(this).html(ordre);
                $.ajax({
                    url: $('#sort').attr('data-action'),
                    method: 'POST',
                    headers: {
                        'X-CSRF-TOKEN': $('meta[name="csrf-token"]').attr('content')
                    },
                    data: {
                        'service': parseInt($(this).prevAll('td.id').html()),
                        'ordre': ordre,
                        '_method': 'PUT'
                    }
                }).done(function (data) {
                        if (console && console.log) {
                            console.dir(data);
                        }
                }); 
            });
        };

    $("#sort tbody").sortable({
        helper: fixHelperModified,
        stop: updateIndex
    }).disableSelection();

    
    $('#filter-form input').change(function(){
        $('#filter-form').submit();
    });
    
    $('a[data-toggle="tab"]').on('shown.bs.tab', function (e) {
        updatePreview();
    });
    
    $('#code').colorpicker({format: 'hex'});
    
    $('.custom-colorpicker button').click(function(e){
        e.preventDefault();
        $(this).parents('.form-group').find('input').val(this.value);
        updatePreview();
    });
    
    $('.activate-input').change(function(e){
        $('.form-control[name="'+e.target.dataset.target+'"]').prop('disabled', !this.checked);
    });
    if(window.tinymce) {
        tinymce.init(
            {
                selector: '.tinymce',
                min_height: 300,
                plugins: [
                    'advlist autolink lists link image charmap print preview hr anchor pagebreak',
                    'searchreplace wordcount visualblocks visualchars code fullscreen',
                    'insertdatetime media nonbreaking save table contextmenu directionality',
                    'emoticons template paste textcolor colorpicker textpattern imagetools'
                ],
                toolbar1: 'insertfile undo redo | styleselect | bold italic | alignleft aligncenter alignright alignjustify | bullist numlist outdent indent | link image media',
                toolbar2: 'forecolor backcolor | preview code',
                image_advtab: true,
                language_url: '/js/tinymce/fr_FR.js',
                content_css: "/css/tinymce.css",
            }
        );
    }
    
    $('input[name="debut"], input[name="fin"]').daterangepicker({
        singleDatePicker: true, 
        locale: {
            format: 'YYYY-MM-DD'
        }
    });
    
    if($('.modal-campaign').length > 0) {
        
        $('.modal-campaign').modal({
            keyboard: false,
            backdrop: 'static'
        });
        
        $('.modal-campaign button').click(function(e){
            var element = e.target;
            $.ajax({
                url: $(this).attr('data-action'),
                method: 'POST',
                headers: {
                    'X-CSRF-TOKEN': $('meta[name="csrf-token"]').attr('content')
                },
                data: {
                    'value': $(this).attr('data-value'),
                    '_method': 'PUT'
                }
            }).done(function (data) {
                $(element).parents('.modal-campaign').modal('hide');
            });
            $(element).parents('.modal-campaign').modal('hide');
        });
        
        // Fix multiple modal scrolling issue
        $('.modal-campaign').on('hidden.bs.modal', function () {
            if($('.modal-campaign').hasClass('in')) {
                $('body').addClass('modal-open');
            }
        });
        
        $('.modal-backdrop').css('opacity', '0');
        $('.modal-backdrop').first().css('opacity', '0.5');
    }
});