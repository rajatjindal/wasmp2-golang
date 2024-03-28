package main

import (
	"net/http"

	spinhttp "github.com/rajatjindal/wasip2-golang/http"
	"github.com/rajatjindal/wasip2-golang/http_trigger"
)

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		spinhttp.NewRouter().GET("/whatever", func(w http.ResponseWriter, r *http.Request, x spinhttp.Params) {
			w.Write([]byte("inside "))
		})

		_ = http_trigger.WasiIo0_2_0_StreamsStreamError{}
		_ = http_trigger.WasiHttp0_2_0_TypesMethod{}
		_ = http_trigger.WasiHttp0_2_0_TypesScheme{}
		_ = http_trigger.WasiHttp0_2_0_TypesDnsErrorPayload{}
		_ = http_trigger.WasiHttp0_2_0_TypesTlsAlertReceivedPayload{}
		_ = http_trigger.WasiHttp0_2_0_TypesFieldSizePayload{}
		_ = http_trigger.WasiHttp0_2_0_TypesErrorCode{}
		_ = http_trigger.WasiHttp0_2_0_TypesHeaderError{}
		_ = http_trigger.WasiHttp0_2_0_TypesTuple2FieldKeyFieldValueT{}
		_ = http_trigger.FermyonSpin2_0_0_LlmInferencingParams{}
		_ = http_trigger.FermyonSpin2_0_0_LlmError{}
		_ = http_trigger.FermyonSpin2_0_0_LlmInferencingUsage{}
		_ = http_trigger.FermyonSpin2_0_0_LlmInferencingResult{}
		_ = http_trigger.FermyonSpin2_0_0_LlmEmbeddingsUsage{}
		_ = http_trigger.FermyonSpin2_0_0_LlmEmbeddingsResult{}
		_ = http_trigger.FermyonSpin2_0_0_RedisError{}
		_ = http_trigger.FermyonSpin2_0_0_RedisRedisParameter{}
		_ = http_trigger.FermyonSpin2_0_0_RedisRedisResult{}
		_ = http_trigger.FermyonSpin2_0_0_MqttError{}
		_ = http_trigger.FermyonSpin2_0_0_MqttQos{}
		_ = http_trigger.FermyonSpin2_0_0_RdbmsTypesError{}
		_ = http_trigger.FermyonSpin2_0_0_RdbmsTypesDbDataType{}
		_ = http_trigger.FermyonSpin2_0_0_RdbmsTypesDbValue{}
		_ = http_trigger.FermyonSpin2_0_0_RdbmsTypesParameterValue{}
		_ = http_trigger.FermyonSpin2_0_0_RdbmsTypesColumn{}
		_ = http_trigger.FermyonSpin2_0_0_RdbmsTypesRowSet{}
		_ = http_trigger.FermyonSpin2_0_0_SqliteError{}
		_ = http_trigger.FermyonSpin2_0_0_SqliteValue{}
		_ = http_trigger.FermyonSpin2_0_0_SqliteRowResult{}
		_ = http_trigger.FermyonSpin2_0_0_SqliteQueryResult{}
		_ = http_trigger.FermyonSpin2_0_0_KeyValueError{}
		_ = http_trigger.FermyonSpin2_0_0_VariablesError{}
		_ = http_trigger.WasiCli0_2_0_EnvironmentTuple2StringStringT{}
		_ = http_trigger.WasiClocks0_2_0_WallClockDatetime{}
		_ = http_trigger.WasiFilesystem0_2_0_TypesDescriptorType{}
		_ = http_trigger.WasiFilesystem0_2_0_TypesDescriptorStat{}
		_ = http_trigger.WasiFilesystem0_2_0_TypesNewTimestamp{}
		_ = http_trigger.WasiFilesystem0_2_0_TypesDirectoryEntry{}
		_ = http_trigger.WasiFilesystem0_2_0_TypesErrorCode{}
		_ = http_trigger.WasiFilesystem0_2_0_TypesAdvice{}
		_ = http_trigger.WasiFilesystem0_2_0_TypesMetadataHashValue{}
		_ = http_trigger.WasiFilesystem0_2_0_TypesTuple2ListU8TBoolT{}
		_ = http_trigger.WasiFilesystem0_2_0_PreopensTuple2DescriptorStringT{}
		_ = http_trigger.WasiSockets0_2_0_NetworkErrorCode{}
		_ = http_trigger.WasiSockets0_2_0_NetworkIpAddressFamily{}
		_ = http_trigger.WasiSockets0_2_0_NetworkIpv4Address{}
		_ = http_trigger.WasiSockets0_2_0_NetworkIpv6Address{}
		_ = http_trigger.WasiSockets0_2_0_NetworkIpAddress{}
		_ = http_trigger.WasiSockets0_2_0_NetworkIpv4SocketAddress{}
		_ = http_trigger.WasiSockets0_2_0_NetworkIpv6SocketAddress{}
		_ = http_trigger.WasiSockets0_2_0_NetworkIpSocketAddress{}
		_ = http_trigger.WasiSockets0_2_0_UdpIncomingDatagram{}
		_ = http_trigger.WasiSockets0_2_0_UdpOutgoingDatagram{}
		_ = http_trigger.WasiSockets0_2_0_UdpTuple2IncomingDatagramStreamOutgoingDatagramStreamT{}
		_ = http_trigger.WasiSockets0_2_0_TcpShutdownType{}
		_ = http_trigger.WasiSockets0_2_0_TcpTuple2InputStreamOutputStreamT{}
		_ = http_trigger.WasiSockets0_2_0_TcpTuple3TcpSocketInputStreamOutputStreamT{}
		_ = http_trigger.WasiRandom0_2_0_InsecureSeedTuple2U64U64T{}

		http_trigger.ExportsWasiHttp0_2_0_IncomingHandler
		w.Write([]byte("OK"))
	})
}

func main() {}
