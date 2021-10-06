import json
from urllib.parse import urlparse,parse_qs
from argparse import Namespace
from pyjfuzz.lib import *

file = "traffic.json"


def mergefile():
    file1 = "traffic-new1.json"
    file2 = "traffic-new2.json"
    file3 = "traffic-new3.json"
    filelist = [file1,file2,file3]

    with open(file,"w") as f:
        for jsonfile in filelist:
            with open(jsonfile,"r") as fff:
                for l in fff:
                    f.writelines(l)

'''
# 测试三种类型url的解析结果
test1 = "/nnrf-disc/v1/nf-instances?requester-nf-type=AMF&target-nf-type=AUSF"
test2 = "/nudr-dr/v1/subscription-data/imsi-208930000000019/20893/provisioned-data/am-data?supported-features="
test3 = "/nudr-dr/v1/subscription-data/imsi-208930000000019/20893/provisioned-data/sm-data?single-nssai={\"sst\":1,\"sd\":\"010203\"}"

demo:
test1 = "/nnrf-disc/v1/nf-instances?requester-nf-type=AMF&target-nf-type=AUSF"
parsed_result = urlparse(test1)
print(parsed_result)
r = parse_qs(parsed_result.query)
print(r)

printout:
test1：{'requester-nf-type': ['AMF'], 'target-nf-type': ['AUSF']}
test2：{}
test3：{'single-nssai': ['{"sst":1,"sd":"010203"}']}
'''


def get_data() -> dict:
    '''
    读出traffic的url并转换成path：qs的格式保存在字典中
    :return:
    '''
    urls = []
    with open(file,"r") as f:
        for l in f:
            data = json.loads(l)
            urls.append(data["Request"]["Url"])
    urls = set(urls)
    ans = {}
    for url in urls:
        parsed_result = urlparse(url)
        path = parsed_result.path
        qs = parse_qs(parsed_result.query)
        if qs!={}:
            ans[path] = qs
    # print(ans)
    return ans

data = get_data()
with open("out.json","w") as f:
    for url,qs in data.items():
        print("before:")
        print(url,qs)
        print("fuzzed:")
        config = PJFConfiguration(Namespace(json=qs, nologo=True, level=6))
        fuzzer = PJFFactory(config)
        i = 0
        while i<10:
            fuzz_data = json.loads(fuzzer.fuzzed)
            qsstr = ""
            for k in fuzz_data:
                for v in fuzz_data[k]:
                    qsstr = qsstr+k+'='+str(v)+'&'
            qsstr=qsstr[:-1]
            print(url+"?"+qsstr)
            out = url+"?"+qsstr
            f.writelines(out+'\n')
            i = i+1