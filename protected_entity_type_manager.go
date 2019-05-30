package arachne

import (
//	"archive/zip"
	"context"
)

type ProtectedEntityTypeManager interface {
   GetTypeName() string
   GetProtectedEntity(ctx context.Context, id ProtectedEntityID) ProtectedEntity
   GetProtectedEntities(ctx context.Context) [] ProtectedEntity
   //Serialize(pe ProtectedEntity, out Zip.Writer)
   //Deserialize(is ZipInputStream, ProtectedEntityInfo peInfo) ProtectedEntity
   //SerializeData(pe ProtectedEntity, out OutputStream)
}
