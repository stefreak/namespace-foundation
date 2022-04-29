// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var languages_nodejs_testdata_services_numberformatter_service_pb = require('../../../../../languages/nodejs/testdata/services/numberformatter/service_pb.js');

function serialize_languages_nodejs_testdata_services_numberformatter_FormatRequest(arg) {
  if (!(arg instanceof languages_nodejs_testdata_services_numberformatter_service_pb.FormatRequest)) {
    throw new Error('Expected argument of type languages.nodejs.testdata.services.numberformatter.FormatRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_languages_nodejs_testdata_services_numberformatter_FormatRequest(buffer_arg) {
  return languages_nodejs_testdata_services_numberformatter_service_pb.FormatRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_languages_nodejs_testdata_services_numberformatter_FormatResponse(arg) {
  if (!(arg instanceof languages_nodejs_testdata_services_numberformatter_service_pb.FormatResponse)) {
    throw new Error('Expected argument of type languages.nodejs.testdata.services.numberformatter.FormatResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_languages_nodejs_testdata_services_numberformatter_FormatResponse(buffer_arg) {
  return languages_nodejs_testdata_services_numberformatter_service_pb.FormatResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var FormatServiceService = exports.FormatServiceService = {
  format: {
    path: '/languages.nodejs.testdata.services.numberformatter.FormatService/Format',
    requestStream: false,
    responseStream: false,
    requestType: languages_nodejs_testdata_services_numberformatter_service_pb.FormatRequest,
    responseType: languages_nodejs_testdata_services_numberformatter_service_pb.FormatResponse,
    requestSerialize: serialize_languages_nodejs_testdata_services_numberformatter_FormatRequest,
    requestDeserialize: deserialize_languages_nodejs_testdata_services_numberformatter_FormatRequest,
    responseSerialize: serialize_languages_nodejs_testdata_services_numberformatter_FormatResponse,
    responseDeserialize: deserialize_languages_nodejs_testdata_services_numberformatter_FormatResponse,
  },
};

exports.FormatServiceClient = grpc.makeGenericClientConstructor(FormatServiceService);
