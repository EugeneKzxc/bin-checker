===============================================================================================================  
[external api request enabled]

data_received..................: 1.3 MB 1.4 kB/s
data_sent......................: 748 kB 805 B/s
dropped_iterations.............: 899985 967.730209/s
http_req_blocked...............: avg=8.15µs   min=0s       med=6µs   max=4.74ms p(90)=9µs   p(95)=12µs 
http_req_connecting............: avg=333ns    min=0s       med=0s    max=398µs  p(90)=0s    p(95)=0s   
http_req_duration..............: avg=1.75s    min=749.1ms  med=1.59s max=4.99s  p(90)=2.39s p(95)=2.66s
 { expected_response:true }...: avg=1.75s    min=749.1ms  med=1.59s max=4.99s  p(90)=2.39s p(95)=2.66s
http_req_failed................: 0.00%  ✓ 0          ✗ 7944
http_req_receiving.............: avg=100.78µs min=5µs      med=92µs  max=3.76ms p(90)=135µs p(95)=178µs
http_req_sending...............: avg=33.7µs   min=2µs      med=30µs  max=1.41ms p(90)=45µs  p(95)=58µs 
http_req_tls_handshaking.......: avg=0s       min=0s       med=0s    max=0s     p(90)=0s    p(95)=0s   
http_req_waiting...............: avg=1.75s    min=748.98ms med=1.59s max=4.99s  p(90)=2.39s p(95)=2.66s
http_reqs......................: 7944   8.541974/s
vus............................: 15     min=15       max=15
vus_max........................: 15     min=15       max=15


===============================================================================================================  
[external api request disabled]

 data_received..................: 2.8 GB   3.1 MB/s
 data_sent......................: 1.6 GB   1.7 MB/s
 dropped_iterations.............: 899968   967.662851/s
 http_req_blocked...............: avg=826ns    min=0s        med=0s     max=40.89ms  p(90)=1µs    p(95)=2µs   
 http_req_connecting............: avg=0ns      min=0s        med=0s     max=983µs    p(90)=0s     p(95)=0s    
 http_req_duration..............: avg=767.5µs  min=69µs      med=439µs  max=144.25ms p(90)=1.31ms p(95)=2.79ms
   { expected_response:true }...: avg=767.5µs  min=69µs      med=439µs  max=144.25ms p(90)=1.31ms p(95)=2.79ms
 http_req_failed................: 0.00%    ✓ 0            ✗ 16811512
 http_req_receiving.............: avg=9.34µs   min=-111000ns med=4µs    max=62.08ms  p(90)=13µs   p(95)=20µs  
 http_req_sending...............: avg=3.53µs   min=0s        med=1µs    max=52.71ms  p(90)=4µs    p(95)=6µs   
 http_req_tls_handshaking.......: avg=0s       min=0s        med=0s     max=0s       p(90)=0s     p(95)=0s    
 http_req_waiting...............: avg=754.62µs min=62µs      med=432µs  max=144.08ms p(90)=1.28ms p(95)=2.74ms
 http_reqs......................: 16811512 18076.060073/s
 iteration_duration.............: avg=11m43s   min=11m43s    med=11m43s max=11m43s   p(90)=11m43s p(95)=11m43s
 iterations.....................: 15       0.016128/s
 vus............................: 2        min=2          max=15    
 vus_max........................: 15       min=15         max=15    
