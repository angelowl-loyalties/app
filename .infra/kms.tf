resource "aws_kms_alias" "angelowl_kms" {
  name          = "alias/angel-owl-kms-key"
  target_key_id = aws_kms_key.angelowl_kms.key_id
}

resource "aws_kms_key" "angelowl_kms" {
  description             = "KMS Key for AngelOwl Loyalty Application"
  deletion_window_in_days = 7
  policy = jsonencode({
    "Id" : "key-consolepolicy-3",
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "Enable IAM User Permissions",
        "Effect" : "Allow",
        "Principal" : {
          "AWS" : "arn:aws:iam::276374573009:root"
        },
        "Action" : "kms:*",
        "Resource" : "*"
      },
      {
        "Sid" : "Allow access for Key Administrators",
        "Effect" : "Allow",
        "Principal" : {
          "AWS" : [
            "arn:aws:iam::276374573009:user/wenlianggg",
            "arn:aws:iam::276374573009:user/itsag1t2"
          ]
        },
        "Action" : [
          "kms:Create*",
          "kms:Describe*",
          "kms:Enable*",
          "kms:List*",
          "kms:Put*",
          "kms:Update*",
          "kms:Revoke*",
          "kms:Disable*",
          "kms:Get*",
          "kms:Delete*",
          "kms:TagResource",
          "kms:UntagResource",
          "kms:ScheduleKeyDeletion",
          "kms:CancelKeyDeletion"
        ],
        "Resource" : "*"
      },
      {
        "Sid" : "Allow use of the key",
        "Effect" : "Allow",
        "Principal" : {
          "AWS" : [
            "arn:aws:iam::276374573009:role/angel-owl-eks-role",
            "arn:aws:iam::276374573009:user/wenlianggg",
            "arn:aws:iam::276374573009:user/alvin_owyong",
            "arn:aws:iam::276374573009:user/awwkl",
            "arn:aws:iam::276374573009:user/justinnnwashere",
            "arn:aws:iam::276374573009:user/lyejianyi",
            "arn:aws:iam::276374573009:user/omerwyo",
            "arn:aws:iam::276374573009:user/oversparkling"
          ]
        },
        "Action" : [
          "kms:Encrypt",
          "kms:Decrypt",
          "kms:ReEncrypt*",
          "kms:GenerateDataKey*",
          "kms:DescribeKey"
        ],
        "Resource" : "*"
      },
      {
        "Sid" : "Allow attachment of persistent resources",
        "Effect" : "Allow",
        "Principal" : {
          "AWS" : [
            "arn:aws:iam::276374573009:role/angel-owl-eks-role",
            "arn:aws:iam::276374573009:user/wenlianggg",
            "arn:aws:iam::276374573009:user/alvin_owyong",
            "arn:aws:iam::276374573009:user/awwkl",
            "arn:aws:iam::276374573009:user/justinnnwashere",
            "arn:aws:iam::276374573009:user/lyejianyi",
            "arn:aws:iam::276374573009:user/omerwyo",
            "arn:aws:iam::276374573009:user/oversparkling"
          ]
        },
        "Action" : [
          "kms:CreateGrant",
          "kms:ListGrants",
          "kms:RevokeGrant"
        ],
        "Resource" : "*",
        "Condition" : {
          "Bool" : {
            "kms:GrantIsForAWSResource" : "true"
          }
        }
      }
    ]
  })
}
