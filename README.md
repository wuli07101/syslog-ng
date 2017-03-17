# syslog-ng 
------
说明：syslog-ng consumer

###PHP产生syslog-ng日志示例代码：

```php
<?php

  
   $data = ;
   $params = array(
        'type'=>'http',
        'msg'  => json_encode(array("fsdfd"=>"fsdfdsf")),
        'url' => "http://www.baidu.com",
   );


 $ident = '/test-syslog/';
 openlog($ident, LOG_PID , LOG_LOCAL6);
 syslog(LOG_INFO, json_encode($params));
 closelog();
```