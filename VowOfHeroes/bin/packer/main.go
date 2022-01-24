package main

import (
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
	pb "github.com/xhebox/revcol/VowOfHeroes/proto"
	json "google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func protoCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proto",
		Short: "decode/encode proto config",
	}
	cmd.AddCommand(&cobra.Command{
		Use:   "ser [json] [asset]",
		Short: "serialize from json to asset",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			in, err := ioutil.ReadFile(args[0])
			if err != nil {
				log.Fatalf("can not read file: %+v\n", err)
			}
			strings := &pb.ProtoConfig{}
			if err := json.Unmarshal(in, strings); err != nil {
				log.Fatalf("failed to parse json: %+v\n", err)
			}
			out, err := proto.Marshal(strings)
			if err != nil {
				log.Fatalf("failed to marshal proto: %+v\n", err)
			}
			if err := ioutil.WriteFile(args[1], out, 0644); err != nil {
				log.Fatalf("can not write file: %+v\n", err)
			}
		},
	})
	cmd.AddCommand(&cobra.Command{
		Use:   "des [asset] [json]",
		Short: "deserialize asset to json",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			in, err := ioutil.ReadFile(args[0])
			if err != nil {
				log.Fatalf("can not read file: %+v\n", err)
			}
			strings := &pb.ProtoConfig{}
			if err := proto.Unmarshal(in, strings); err != nil {
				log.Fatalf("failed to parse proto: %+v\n", err)
			}
			marshaler := json.MarshalOptions{
				Multiline: true,
				Indent: "\t",
			}
			out, err := marshaler.Marshal(strings)
			if err != nil {
				log.Fatalf("failed to marshal json: %+v\n", err)
			}
			if err := ioutil.WriteFile(args[1], out, 0644); err != nil {
				log.Fatalf("can not write file: %+v\n", err)
			}
		},
	})
	return cmd
}

func strResCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "str_res",
		Short: "decode/encode StrRes asset",
	}
	cmd.AddCommand(&cobra.Command{
		Use:   "ser [json] [asset]",
		Short: "serialize from json to asset",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			in, err := ioutil.ReadFile(args[0])
			if err != nil {
				log.Fatalf("can not read file: %+v\n", err)
			}
			strings := &pb.StrResAsset{}
			if err := json.Unmarshal(in, strings); err != nil {
				log.Fatalf("failed to parse json: %+v\n", err)
			}
			out, err := proto.Marshal(strings)
			if err != nil {
				log.Fatalf("failed to marshal proto: %+v\n", err)
			}
			if err := ioutil.WriteFile(args[1], out, 0644); err != nil {
				log.Fatalf("can not write file: %+v\n", err)
			}
		},
	})
	cmd.AddCommand(&cobra.Command{
		Use:   "des [asset] [json]",
		Short: "deserialize asset to json",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			in, err := ioutil.ReadFile(args[0])
			if err != nil {
				log.Fatalf("can not read file: %+v\n", err)
			}
			strings := &pb.StrResAsset{}
			if err := proto.Unmarshal(in, strings); err != nil {
				log.Fatalf("failed to parse proto: %+v\n", err)
			}
			marshaler := json.MarshalOptions{
				Multiline: true,
				Indent: "\t",
			}
			out, err := marshaler.Marshal(strings)
			if err != nil {
				log.Fatalf("failed to marshal json: %+v\n", err)
			}
			if err := ioutil.WriteFile(args[1], out, 0644); err != nil {
				log.Fatalf("can not write file: %+v\n", err)
			}
		},
	})
	return cmd
}

func attributeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "attribute",
		Short: "decode/encode attribute asset",
	}
	cmd.AddCommand(&cobra.Command{
		Use:   "ser [json] [asset]",
		Short: "serialize from json to asset",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			in, err := ioutil.ReadFile(args[0])
			if err != nil {
				log.Fatalf("can not read file: %+v\n", err)
			}
			strings := &pb.AttributeAsset{}
			if err := json.Unmarshal(in, strings); err != nil {
				log.Fatalf("failed to parse json: %+v\n", err)
			}
			out, err := proto.Marshal(strings)
			if err != nil {
				log.Fatalf("failed to marshal proto: %+v\n", err)
			}
			if err := ioutil.WriteFile(args[1], out, 0644); err != nil {
				log.Fatalf("can not write file: %+v\n", err)
			}
		},
	})
	cmd.AddCommand(&cobra.Command{
		Use:   "des [asset] [json]",
		Short: "deserialize asset to json",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			in, err := ioutil.ReadFile(args[0])
			if err != nil {
				log.Fatalf("can not read file: %+v\n", err)
			}
			strings := &pb.AttributeAsset{}
			if err := proto.Unmarshal(in, strings); err != nil {
				log.Fatalf("failed to parse proto: %+v\n", err)
			}
			marshaler := json.MarshalOptions{
				Multiline: true,
				Indent: "\t",
			}
			out, err := marshaler.Marshal(strings)
			if err != nil {
				log.Fatalf("failed to marshal json: %+v\n", err)
			}
			if err := ioutil.WriteFile(args[1], out, 0644); err != nil {
				log.Fatalf("can not write file: %+v\n", err)
			}
		},
	})
	return cmd
}

func monsterCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "monster",
		Short: "decode/encode monster asset",
	}
	cmd.AddCommand(&cobra.Command{
		Use:   "ser [json] [asset]",
		Short: "serialize from json to asset",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			in, err := ioutil.ReadFile(args[0])
			if err != nil {
				log.Fatalf("can not read file: %+v\n", err)
			}
			strings := &pb.MonsterAsset{}
			if err := json.Unmarshal(in, strings); err != nil {
				log.Fatalf("failed to parse json: %+v\n", err)
			}
			out, err := proto.Marshal(strings)
			if err != nil {
				log.Fatalf("failed to marshal proto: %+v\n", err)
			}
			if err := ioutil.WriteFile(args[1], out, 0644); err != nil {
				log.Fatalf("can not write file: %+v\n", err)
			}
		},
	})
	cmd.AddCommand(&cobra.Command{
		Use:   "des [asset] [json]",
		Short: "deserialize asset to json",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			in, err := ioutil.ReadFile(args[0])
			if err != nil {
				log.Fatalf("can not read file: %+v\n", err)
			}
			strings := &pb.MonsterAsset{}
			if err := proto.Unmarshal(in, strings); err != nil {
				log.Fatalf("failed to parse proto: %+v\n", err)
			}
			marshaler := json.MarshalOptions{
				Multiline: true,
				Indent: "\t",
			}
			out, err := marshaler.Marshal(strings)
			if err != nil {
				log.Fatalf("failed to marshal json: %+v\n", err)
			}
			if err := ioutil.WriteFile(args[1], out, 0644); err != nil {
				log.Fatalf("can not write file: %+v\n", err)
			}
		},
	})
	return cmd
}

func monsterGroupCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "monster_group",
		Short: "decode/encode attribute asset",
	}
	cmd.AddCommand(&cobra.Command{
		Use:   "ser [json] [asset]",
		Short: "serialize from json to asset",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			in, err := ioutil.ReadFile(args[0])
			if err != nil {
				log.Fatalf("can not read file: %+v\n", err)
			}
			strings := &pb.MonsterGroupAsset{}
			if err := json.Unmarshal(in, strings); err != nil {
				log.Fatalf("failed to parse json: %+v\n", err)
			}
			out, err := proto.Marshal(strings)
			if err != nil {
				log.Fatalf("failed to marshal proto: %+v\n", err)
			}
			if err := ioutil.WriteFile(args[1], out, 0644); err != nil {
				log.Fatalf("can not write file: %+v\n", err)
			}
		},
	})
	cmd.AddCommand(&cobra.Command{
		Use:   "des [asset] [json]",
		Short: "deserialize asset to json",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			in, err := ioutil.ReadFile(args[0])
			if err != nil {
				log.Fatalf("can not read file: %+v\n", err)
			}
			strings := &pb.MonsterGroupAsset{}
			if err := proto.Unmarshal(in, strings); err != nil {
				log.Fatalf("failed to parse proto: %+v\n", err)
			}
			marshaler := json.MarshalOptions{
				Multiline: true,
				Indent: "\t",
			}
			out, err := marshaler.Marshal(strings)
			if err != nil {
				log.Fatalf("failed to marshal json: %+v\n", err)
			}
			if err := ioutil.WriteFile(args[1], out, 0644); err != nil {
				log.Fatalf("can not write file: %+v\n", err)
			}
		},
	})
	return cmd
}

func main() {
	var rootCmd = &cobra.Command{
		Short: "Serializer and Deserializer for text assets",
	}
	rootCmd.AddCommand(protoCommand())
	rootCmd.AddCommand(strResCommand())
	rootCmd.AddCommand(attributeCommand())
	rootCmd.AddCommand(monsterCommand())
	rootCmd.AddCommand(monsterGroupCommand())
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("%+v\n", err)
	}
}
